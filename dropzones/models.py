from typing import TYPE_CHECKING, Optional
from sqlmodel import Relationship, SQLModel, Field

if TYPE_CHECKING:
    from columns.models import Column


class Dropzone(SQLModel, table=True):
    id: Optional[int] = Field(None, primary_key=True)
    name: str
    order: int
    column_id: int
    column: Optional["Column"] = Relationship(back_populates="dropzones")
