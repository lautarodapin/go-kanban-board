from typing import List
from fastapi import APIRouter, HTTPException
from sqlmodel import Session, select
from sqlmodel.ext.asyncio.session import AsyncSession
from models.database import engine

from .models import Board, CreateBoard

api = APIRouter(prefix='/boards')


@api.get('/', response_model=list[Board])
async def get_boards():
    with Session(engine) as session:
        results = session.exec(select(Board)).fetchall()
    return results


@api.get('/{board_id}', response_model=Board)
async def get_board(board_id: int):
    with Session(engine) as session:
        result = session.exec(select(Board).where(
            Board.id == board_id)).first()
        if result:
            return result
    raise HTTPException(404, 'Board not found')


@api.post('/', response_model=Board)
async def create_board(board: CreateBoard):
    with Session(engine) as session:
        board_db = Board.from_orm(board)
        session.add(board_db)
        session.commit()
        session.refresh(board_db)
    return board_db


@api.put('/{board_id}', response_model=Board)
async def update_board(board_id: int, board: Board):
    with Session(engine) as session:
        board_db = session.get(Board, board_id)
        if board_db is None:
            raise HTTPException(404, 'Board not found')

        for k, v in board.dict(exclude_unset=True).items():
            setattr(board_db, k, v)
        session.add(board_db)
        session.commit()
        session.refresh(board_db)
    return board_db


@api.delete('/{board_id}')
async def delete_board(board_id: int):
    with Session(engine) as session:
        board = session.exec(select(Board).where(Board.id == board_id)).one()
        session.delete(board)
        session.commit()
