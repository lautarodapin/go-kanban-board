from typing import TYPE_CHECKING, Optional
from sqlmodel import Relationship, SQLModel, Field


if TYPE_CHECKING:
    pass


class Column(SQLModel, table=True):
    id: Optional[int] = Field(None, primary_key=True)
    name: str
    order: int
