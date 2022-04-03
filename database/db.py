from sqlmodel import SQLModel, Session, create_engine
from sqlalchemy.ext.asyncio import create_async_engine
from sqlmodel.ext.asyncio.session import AsyncSession


sqlite_file_name = "database.sqlite"
sqlite_url = f"sqlite:///{sqlite_file_name}"
async_sqlite_url = f"sqlite+aiosqlite:///{sqlite_file_name}"

engine = create_engine(sqlite_url, echo=True)
async_engine = create_async_engine(async_sqlite_url, echo=True)


def create_db_and_tables():
    from boards.models import Board
    from dropzones.models import Dropzone
    from columns.models import Column
    from tickets.models import Ticket

    SQLModel.metadata.create_all(engine)


async def get_async_session():
    async with AsyncSession(async_engine) as session:
        yield session


def get_session():
    with Session(engine) as session:
        yield session
