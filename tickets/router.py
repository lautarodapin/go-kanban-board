from fastapi import APIRouter
from sqlmodel import Session, select

from models.database import engine
from .models import Ticket

api = APIRouter(prefix='/tickets')


@api.get('/', response_model=list[Ticket])
async def get_tickets():
    with Session(engine) as session:
        results = session.exec(select(Ticket))
    return results


@api.get('/{ticket_id}', response_model=Ticket)
async def get_ticket(ticket_id: int):
    with Session(engine) as session:
        result = session.exec(select(Ticket).where(Ticket.id == ticket_id))
    return result[0]


@api.post('/', response_model=Ticket)
async def create_ticket(ticket: Ticket):
    with Session(engine) as session:
        session.add(ticket)
        await session.commit()
    return ticket


@api.put('/{ticket_id}', response_model=Ticket)
async def update_ticket(ticket_id: int, ticket: Ticket):
    with Session(engine) as session:
        session.add
        session.update(ticket, where=Ticket.id == ticket_id)
    return ticket


@api.delete('/{ticket_id}')
async def delete_ticket(ticket_id: int):
    with Session(engine) as session:
        session.delete(Ticket, where=Ticket.id == ticket_id)
