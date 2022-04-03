from typing import TYPE_CHECKING, Optional
from sqlmodel import Relationship, SQLModel, Field


if TYPE_CHECKING:
    from boards.models import Board
    from dropzones.models import Dropzone


class BaseTicket(SQLModel):
    title: str
    description: Optional[str]
    dropzone_id: Optional[int] = Field(None, foreign_key='dropzone.id')
    board_id: Optional[int] = Field(None, foreign_key='board.id')


class Ticket(BaseTicket, table=True):
    id: Optional[int] = Field(None, primary_key=True)

    dropzone: Optional["Dropzone"] = Relationship(back_populates="tickets")
    board: Optional["Board"] = Relationship(back_populates="tickets")
