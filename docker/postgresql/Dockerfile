FROM postgres:12-alpine

VOLUME $PWD/db:/docker-entrypoint-initdb.d

RUN echo "listen_addresses = 'web'" >> /var/lib/postgresql/data/postgresql.conf