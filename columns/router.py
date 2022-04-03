from fastapi import APIRouter, HTTPException
from sqlmodel import select, Session
from database.db import engine

from .models import Column


api = APIRouter(prefix='/columns')


@api.get('/', response_model=list[Column])
async def get_columns():
    with Session(engine) as session:
        results = session.exec(select(Column)).fetchall()
    return results


@api.get('/{column_id}', response_model=Column)
async def get_column(column_id: int):
    with Session(engine) as session:
        result = session.exec(select(Column).where(
            Column.id == column_id)).first()
    if result:
        return result
    raise HTTPException(404, 'Column not found')


@api.post('/', response_model=Column)
async def create_column(column: Column):
    with Session(engine) as session:
        column_db = Column.from_orm(column)
        session.add(column_db)
        session.commit()
        session.refresh(column_db)
    return column_db


@api.put('/{column_id}', response_model=Column)
async def update_column(column_id: int, column: Column):
    with Session(engine) as session:
        column_db = session.get(Column, column_id)
        if column_db is None:
            raise HTTPException(404, 'Column not found')

        for k, v in column.dict(exclude_unset=True).items():
            setattr(column_db, k, v)
        session.add(column_db)
        session.commit()
        session.refresh(column_db)
    return column_db


@api.delete('/{column_id}')
async def delete_column(column_id: int):
    with Session(engine) as session:
        session.delete(Column, where=Column.id == column_id)
        session.commit()
