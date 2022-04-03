from fastapi import APIRouter, Depends, HTTPException
from sqlmodel import select, Session
from database.db import engine, get_async_session, AsyncSession
from .models import CreateDropzone, Dropzone, UpdateDropzone


api = APIRouter(prefix='/dropzones')


@api.get('/', response_model=list[Dropzone])
async def get_dropzones(
    session: AsyncSession = Depends(get_async_session),
):
    return (await session.exec(select(Dropzone))).fetchall()


@api.get('/{dropzone_id}', response_model=Dropzone)
async def get_dropzone(
    dropzone_id: int,
    session: AsyncSession = Depends(get_async_session),
):
    result = await session.get(Dropzone, dropzone_id)
    if result:
        return result
    raise HTTPException(404, 'Dropzone not found')


@api.post('/', response_model=Dropzone)
async def create_dropzone(
    dropzone: CreateDropzone,
    session: AsyncSession = Depends(get_async_session),
):
    dropzone_db = Dropzone.from_orm(dropzone)
    session.add(dropzone_db)
    await session.commit()
    await session.refresh(dropzone_db)
    return dropzone_db


@api.put('/{dropzone_id}', response_model=Dropzone)
async def update_dropzone(
    dropzone_id: int,
    dropzone: UpdateDropzone,
    session: AsyncSession = Depends(get_async_session),
):
    dropzone_db = await session.get(Dropzone, dropzone_id)
    if dropzone_db is None:
        raise HTTPException(404, 'Dropzone not found')

    for k, v in dropzone.dict(exclude_unset=True).items():
        setattr(dropzone_db, k, v)
    session.add(dropzone_db)
    await session.commit()
    await session.refresh(dropzone_db)
    return dropzone_db


@api.delete('/{dropzone_id}')
async def delete_dropzone(
    dropzone_id: int,
    session: AsyncSession = Depends(get_async_session),
):
    dropzone_db = await session.get(Dropzone, dropzone_id)
    await session.delete(dropzone_db)
    await session.commit()
