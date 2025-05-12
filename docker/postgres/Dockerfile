FROM postgres:16

# Copy initialization script
COPY ../../db/init.sql /docker-entrypoint-initdb.d/

# Environment variables can also be passed via docker-compose
