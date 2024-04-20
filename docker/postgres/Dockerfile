FROM postgres:latest

# Set environment variables
ENV POSTGRES_USER=admin
ENV POSTGRES_PASSWORD=admin
ENV POSTGRES_DB=go_chi_db

COPY init.sql /docker-entrypoint-initdb.d/
