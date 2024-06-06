FROM python:3.12

WORKDIR /app

COPY /pyproject.toml /app

RUN pip3 install poetry

RUN poetry config virtualenvs.create false

RUN poetry install 

COPY . /app

CMD [ "python", "server.py" ] 