from typing import TYPE_CHECKING, List, Optional
from sqlmodel import Field, Relationship, SQLModel


if TYPE_CHECKING:
    from tickets.models import Ticket


class BaseBoard(SQLModel):
    name: str
    description: Optional[str]


class Board(BaseBoard, table=True):
    id: Optional[int] = Field(None, primary_key=True)
    tickets: List["Ticket"] = Relationship(back_populates="board")


class CreateBoard(BaseBoard):
    ...


class UpdateBoard(BaseBoard):
    ...
