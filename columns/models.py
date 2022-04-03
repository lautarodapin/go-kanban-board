from typing import TYPE_CHECKING, List, Optional
from sqlmodel import Relationship, SQLModel, Field


if TYPE_CHECKING:
    from dropzones.models import Dropzone


class BaseColumn(SQLModel):
    name: str
    order: int


class Column(BaseColumn, table=True):
    id: Optional[int] = Field(None, primary_key=True)
    dropzones: List["Dropzone"] = Relationship(back_populates="column")


class CreateColumn(BaseColumn):
    ...


class UpdateColumn(BaseColumn):
    ...
