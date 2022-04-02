from fastapi import APIRouter
from sqlmodel import select, Session
from models.database import engine

from .models import Column


api = APIRouter(prefix='/columns')


@api.get('/', response_model=list[Column])
async def get_columns():
    with Session(engine) as session:
        results = session.exec(select(Column))
    return results


@api.get('/{column_id}', response_model=Column)
async def get_column(column_id: int):
    with Session(engine) as session:
        result = session.exec(select(Column).where(Column.id == column_id))
    return result[0]


@api.post('/', response_model=Column)
async def create_column(column: Column):
    with Session(engine) as session:
        session.add(column)
        await session.commit()
    return column


@api.put('/{column_id}', response_model=Column)
async def update_column(column_id: int, column: Column):
    with Session(engine) as session:
        session.update(column, where=Column.id == column_id)
    return column


@api.delete('/{column_id}')
async def delete_column(column_id: int):
    with Session(engine) as session:
        session.delete(Column, where=Column.id == column_id)
        session.commit()
