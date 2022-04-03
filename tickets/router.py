from fastapi import APIRouter, HTTPException
from sqlmodel import Session, select

from models.database import engine
from .models import Ticket

api = APIRouter(prefix='/tickets')


@api.get('/', response_model=list[Ticket])
async def get_tickets():
    with Session(engine) as session:
        results = session.exec(select(Ticket)).fetchall()
    return results


@api.get('/{ticket_id}', response_model=Ticket)
async def get_ticket(ticket_id: int):
    with Session(engine) as session:
        result = session.exec(select(Ticket).where(
            Ticket.id == ticket_id)).first()
    if result:
        return result
    raise HTTPException(404, 'Ticket not found')


@api.post('/', response_model=Ticket)
async def create_ticket(ticket: Ticket):
    with Session(engine) as session:
        ticket_db = Ticket.from_orm(ticket)
        session.add(ticket_db)
        session.commit()
        session.refresh(ticket_db)
    return ticket_db


@api.put('/{ticket_id}', response_model=Ticket)
async def update_ticket(ticket_id: int, ticket: Ticket):
    with Session(engine) as session:
        ticket_db = session.get(Ticket, ticket_id)
        if ticket_db is None:
            raise HTTPException(404, 'Ticket not found')

        for k, v in ticket.dict(exclude_unset=True).items():
            setattr(ticket_db, k, v)
        session.add(ticket_db)
        session.commit()
        session.refresh(ticket_db)
    return ticket_db


@api.delete('/{ticket_id}')
async def delete_ticket(ticket_id: int):
    with Session(engine) as session:
        session.delete(Ticket, where=Ticket.id == ticket_id)
