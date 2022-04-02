from fastapi import APIRouter
from sqlmodel import select, Session
from models.database import engine
from .models import Dropzone


api = APIRouter(prefix='/dropzones')

@api.get('/', response_model=list[Dropzone])
async def get_dropzones():
    with Session(engine) as session:
        results = session.exec(select(Dropzone))
    return results


@api.get('/{dropzone_id}', response_model=Dropzone)
async def get_dropzone(dropzone_id: int):
    with Session(engine) as session:
        result = session.exec(select(Dropzone).where(Dropzone.id == dropzone_id))
    return result[0]


@api.post('/', response_model=Dropzone)
async def create_dropzone(dropzone: Dropzone):
    with Session(engine) as session:
        session.add(dropzone)
        await session.commit()
    return dropzone


@api.put('/{dropzone_id}', response_model=Dropzone)
async def update_dropzone(dropzone_id: int, dropzone: Dropzone):
    with Session(engine) as session:
        session.add
        session.update(dropzone, where=Dropzone.id == dropzone_id)
    return dropzone


@api.delete('/{dropzone_id}')
async def delete_dropzone(dropzone_id: int):
    with Session(engine) as session:
        session.delete(Dropzone, where=Dropzone.id == dropzone_id)
