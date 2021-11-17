FROM postgres

COPY 000001_init.up.sql /docker-entrypoint-initdb.d/1.sql

CMD ["postgres"]