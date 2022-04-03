from fastapi import APIRouter, HTTPException
from sqlmodel import select, Session
from database.db import engine
from .models import Dropzone


api = APIRouter(prefix='/dropzones')


@api.get('/', response_model=list[Dropzone])
async def get_dropzones():
    with Session(engine) as session:
        results = session.exec(select(Dropzone)).fetchall()
    return results


@api.get('/{dropzone_id}', response_model=Dropzone)
async def get_dropzone(dropzone_id: int):
    with Session(engine) as session:
        result = session.exec(select(Dropzone).where(
            Dropzone.id == dropzone_id)).first()
    if result:
        return result
    raise HTTPException(404, 'Dropzone not found')


@api.post('/', response_model=Dropzone)
async def create_dropzone(dropzone: Dropzone):
    with Session(engine) as session:
        dropzone_db = Dropzone.from_orm(dropzone)
        session.add(dropzone_db)
        session.commit()
        session.refresh(dropzone_db)
    return dropzone_db


@api.put('/{dropzone_id}', response_model=Dropzone)
async def update_dropzone(dropzone_id: int, dropzone: Dropzone):
    with Session(engine) as session:
        dropzone_db = session.get(Dropzone, dropzone_id)
        if dropzone_db is None:
            raise HTTPException(404, 'Dropzone not found')

        for k, v in dropzone.dict(exclude_unset=True).items():
            setattr(dropzone_db, k, v)
        session.add(dropzone_db)
        session.commit()
        session.refresh(dropzone_db)
    return dropzone_db


@api.delete('/{dropzone_id}')
async def delete_dropzone(dropzone_id: int):
    with Session(engine) as session:
        session.delete(Dropzone, where=Dropzone.id == dropzone_id)
