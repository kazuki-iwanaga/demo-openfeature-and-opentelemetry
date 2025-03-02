FROM python:3.12-bookworm

RUN apt-get update \
    && apt-get install -y \
        iputils-ping \
        iproute2 \
        dnsutils

WORKDIR /app

COPY ./requirements.txt .

RUN pip install -r requirements.txt

CMD ["fastapi", "dev", "main.py", "--proxy-headers", "--host", "0.0.0.0", "--port", "8080"]
