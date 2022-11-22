FROM alpine

COPY ./build .

RUN apk update && \
    apk add postgresql-client && \
    apk add --no-cache libc6-compat && \
    chmod +x wait-for-postgres.sh