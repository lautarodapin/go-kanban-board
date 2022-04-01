from typing import Optional

from sqlmodel import Field, SQLModel, create_engine

sqlite_file_name = "database.sqlite"
sqlite_url = f"sqlite:///{sqlite_file_name}"

engine = create_engine(sqlite_url, echo=True)


def create_db_and_tables():
    from boards.models import Board
    from dropzones.models import Dropzone
    from columns.models import Column
    from tickets.models import Ticket

    SQLModel.metadata.create_all(engine)
    
