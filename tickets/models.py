from typing import TYPE_CHECKING, Optional
import enum
from sqlmodel import Relationship, SQLModel, Field, Enum, Column


if TYPE_CHECKING:
    from boards.models import Board
    from dropzones.models import Dropzone


class TicketType(str, enum.Enum):
    bug = "bug"
    feature = "feature"
    task = "task"


class BaseTicket(SQLModel):
    title: str
    description: Optional[str]
    dropzone_id: Optional[int] = Field(None, foreign_key='dropzone.id')
    type: TicketType = Field(sa_column=Column(Enum(TicketType)))


class Ticket(BaseTicket, table=True):
    id: Optional[int] = Field(None, primary_key=True)

    dropzone: Optional["Dropzone"] = Relationship(back_populates="tickets")
    board: Optional["Board"] = Relationship(back_populates="tickets")


class CreateTicket(BaseTicket):
    ...


class UpdateTicket(BaseTicket):
    ...
