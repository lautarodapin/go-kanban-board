from typing import TYPE_CHECKING, List, Optional
from sqlmodel import Relationship, SQLModel, Field


if TYPE_CHECKING:
    from tickets.models import Ticket
    from columns.models import Column


class BaseDropzone(SQLModel):
    name: str
    order: int
    column_id: Optional[int] = Field(None, foreign_key='column.id')


class Dropzone(BaseDropzone, table=True):
    id: Optional[int] = Field(None, primary_key=True)
    column: Optional["Column"] = Relationship(back_populates="dropzones")
    tickets: List["Ticket"] = Relationship(back_populates="dropzone")


class CreateDropzone(BaseDropzone):
    ...


class UpdateDropzone(BaseDropzone):
    ...
