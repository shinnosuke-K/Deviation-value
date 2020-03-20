FROM postgres:11-alpine

RUN echo "listen_addresses = 'web'" >> /var/lib/postgresql/data/postgresql.conf