from typing import TYPE_CHECKING, Optional
from sqlmodel import Relationship, SQLModel, Field


if TYPE_CHECKING:
    from boards.models import Board
    from dropzones.models import Dropzone


class Ticket(SQLModel, table=True):
    id: Optional[int] = Field(None, primary_key=True)
    title: str
    description: Optional[str]

    dropzone_id: int
    dropzone: Optional["Dropzone"] = Relationship(back_populates="tickets")

    board_id: int
    board: Optional["Board"] = Relationship(back_populates="tickets")
