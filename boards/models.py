from typing import TYPE_CHECKING, Optional
from sqlmodel import Field, SQLModel


if TYPE_CHECKING:
    pass


class Board(SQLModel, table=True):
    id: Optional[int] = Field(None, primary_key=True)
    name: str
    description: Optional[str]
