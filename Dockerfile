FROM alpine

COPY ./build .

RUN apk update && \
    apk add postgresql-client && \
    chmod +x wait-for-postgres.sh