from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy import column
from sqlmodel import select, Session
from database.db import engine, get_async_session, AsyncSession

from .models import Column


api = APIRouter(prefix='/columns')


@api.get('/', response_model=list[Column])
async def get_columns(
    session: AsyncSession = Depends(get_async_session),
):
    return (await session.exec(select(Column))).fetchall()


@api.get('/{column_id}', response_model=Column)
async def get_column(
    column_id: int,
    session: AsyncSession = Depends(get_async_session),
):
    result = await session.get(Column, column_id)
    if not result:
        raise HTTPException(404, 'Column not found')
    return result


@api.post('/', response_model=Column)
async def create_column(
    column: Column,
    session: AsyncSession = Depends(get_async_session),
):
    column_db = Column.from_orm(column)
    session.add(column_db)
    await session.commit()
    await session.refresh(column_db)
    return column_db


@api.put('/{column_id}', response_model=Column)
async def update_column(
    column_id: int,
    column: Column,
    session: AsyncSession = Depends(get_async_session),
):
    column_db = await session.get(Column, column_id)
    if column_db is None:
        raise HTTPException(404, 'Column not found')

    for k, v in column.dict(exclude_unset=True).items():
        setattr(column_db, k, v)
    session.add(column_db)
    await session.commit()
    await session.refresh(column_db)
    return column_db


@api.delete('/{column_id}')
async def delete_column(
    column_id: int,
    session: AsyncSession = Depends(get_async_session),
):
    column_db = await session.get(Column, column_id)
    await session.delete(column_db)
    await session.commit()
