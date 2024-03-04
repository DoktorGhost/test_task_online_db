FROM postgres:latest

ENV POSTGRES_USER admin
ENV POSTGRES_PASSWORD admin
ENV POSTGRES_DB testdb

COPY schema.sql /docker-entrypoint-initdb.d/