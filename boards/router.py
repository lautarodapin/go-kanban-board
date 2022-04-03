from typing import List
from fastapi import APIRouter, Depends, HTTPException
from sqlmodel import Session, select
from sqlmodel.ext.asyncio.session import AsyncSession
from database.db import engine, async_engine, get_async_session

from .models import Board, CreateBoard, UpdateBoard

api = APIRouter(prefix='/boards')


@api.get('/', response_model=list[Board])
async def get_boards(session: AsyncSession = Depends(get_async_session)):
    return (await session.exec(select(Board))).fetchall()


@api.get('/{board_id}', response_model=Board)
async def get_board(
    board_id: int,
    session: AsyncSession = Depends(get_async_session),
):
    result = (await session.exec(select(Board).where(
        Board.id == board_id))).first()
    if result:
        return result
    raise HTTPException(404, 'Board not found')


@api.post('/', response_model=Board)
async def create_board(
    board: CreateBoard,
    session: AsyncSession = Depends(get_async_session),
):
    board_db = Board.from_orm(board)
    session.add(board_db)
    await session.commit()
    await session.refresh(board_db)
    return board_db


@api.put('/{board_id}', response_model=Board)
async def update_board(
    board_id: int,
    board: UpdateBoard,
    session: AsyncSession = Depends(get_async_session),
):
    board_db = await session.get(Board, board_id)
    if board_db is None:
        raise HTTPException(404, 'Board not found')

    for k, v in board.dict(exclude_unset=True).items():
        setattr(board_db, k, v)
    session.add(board_db)
    await session.commit()
    await session.refresh(board_db)
    return board_db


@api.delete('/{board_id}')
async def delete_board(
    board_id: int,
    session: AsyncSession = Depends(get_async_session),
):
    board = (await session.exec(select(Board).where(Board.id == board_id))).one()
    await session.delete(board)
    await session.commit()
