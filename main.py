from fastapi import FastAPI
import boards.router
import dropzones.router
import columns.router
from models.database import create_db_and_tables
import tickets.router

def create_app():
    app = FastAPI()
    app.include_router(boards.router.api)
    app.include_router(dropzones.router.api)
    app.include_router(columns.router.api)
    app.include_router(tickets.router.api)
    return app

app = create_app()
create_db_and_tables()