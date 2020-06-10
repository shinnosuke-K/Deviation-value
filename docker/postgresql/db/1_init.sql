CREATE TABLE infos (
    id serial NOT NULL  PRIMARY KEY,
    deviation TEXT,
    school_name TEXT,
    course TEXT,
    url TEXT,
    prefecture TEXT
);


COPY infos FROM '/docker-entrypoint-initdb.d/infos.csv' WITH CSV;