from fastapi import APIRouter
from sqlmodel import Session, select
from models.database import engine

from .models import Board

api = APIRouter(prefix='/boards')


@api.get('/', response_model=list[Board])
async def get_boards():
    with Session(engine) as session:
        results = session.exec(select(Board))
    return results


@api.get('/{board_id}', response_model=Board)
async def get_board(board_id: int):
    with Session(engine) as session:
        result = session.exec(select(Board).where(Board.id == board_id))
    return result[0]


@api.post('/', response_model=Board)
async def create_board(board: Board):
    with Session(engine) as session:
        session.add(board)
        await session.commit()
    return board


@api.put('/{board_id}', response_model=Board)
async def update_board(board_id: int, board: Board):
    with Session(engine) as session:
        board_db = session.exec(select(Board).where(Board.id == board_id))
        board_db = Board(**board.dict())
        board_db.id = board_id
        session.add(board_db)
        session.commit()
    return board_db


@api.delete('/{board_id}')
async def delete_board(board_id: int):
    with Session(engine) as session:
        board = session.exec(select(Board).where(Board.id == board_id)).one()
        session.delete(board)
        session.commit()
