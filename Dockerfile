FROM python:3.10-slim AS python_base

COPY requirements.txt /usr/src/app/requirements.txt

WORKDIR /usr/src/app
RUN pip install -r requirements.txt

COPY . .


FROM python_base AS worker

ENTRYPOINT ["python", "-u","worker.py"]


FROM python_base AS web
ENTRYPOINT ["flask", "run", "--host=0.0.0.0"]