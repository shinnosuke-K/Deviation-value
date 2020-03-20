FROM postgres:11-alpine



# COPY $PWD/docker/postgresql.conf /var/lib/postgresql/data/postgresql.conf
RUN echo "listen_addresses = 'web'" >> /var/lib/postgresql/data/postgresql.conf

# CMD [ "postgres", "-D", "/var/lib/postgresql/data" ]