from fastapi import APIRouter, Depends, HTTPException
from sqlmodel import Session, select

from database.db import engine, get_async_session, AsyncSession
from .models import CreateTicket, Ticket, UpdateTicket

api = APIRouter(prefix='/tickets')


@api.get('/', response_model=list[Ticket])
async def get_tickets(
    session: AsyncSession = Depends(get_async_session),
):
    return (await session.exec(select(Ticket))).fetchall()


@api.get('/{ticket_id}', response_model=Ticket)
async def get_ticket(
    ticket_id: int,
    session: AsyncSession = Depends(get_async_session),
):
    result = await session.get(Ticket, ticket_id)
    if result:
        return result
    raise HTTPException(404, 'Ticket not found')


@api.post('/', response_model=Ticket)
async def create_ticket(
    ticket: CreateTicket,
    session: AsyncSession = Depends(get_async_session),
):
    ticket_db = Ticket.from_orm(ticket)
    await session.add(ticket_db)
    await session.commit()
    await session.refresh(ticket_db)
    return ticket_db


@api.put('/{ticket_id}', response_model=Ticket)
async def update_ticket(
    ticket_id: int,
    ticket: UpdateTicket,
    session: AsyncSession = Depends(get_async_session),
):
    ticket_db = await session.get(Ticket, ticket_id)
    if ticket_db is None:
        raise HTTPException(404, 'Ticket not found')

    for k, v in ticket.dict(exclude_unset=True).items():
        setattr(ticket_db, k, v)
    await session.add(ticket_db)
    await session.commit()
    await session.refresh(ticket_db)
    return ticket_db


@api.delete('/{ticket_id}')
async def delete_ticket(
    ticket_id: int,
    session: AsyncSession = Depends(get_async_session),
):
    ticket_db = await session.get(Ticket, ticket_id)
    await session.delete(ticket_db)
    await session.commit()
