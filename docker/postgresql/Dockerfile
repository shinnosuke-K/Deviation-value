FROM postgres:11-alpine

VOLUME $PWD/docker/postgresql/db:/docker-entrypoint-initdb.d

RUN echo "listen_addresses = 'web'" >> /var/lib/postgresql/data/postgresql.conf