#! /bin/bash -e
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE USER $DB_USER WITH PASSWORD '$DB_PASS';
	CREATE DATABASE $DB_NAME;
    \c $DB_NAME;
    CREATE TABLE containers (
        ip VARCHAR(15) PRIMARY KEY,
        ping SMALLINT CHECK (ping >= 0),
        date DATE
    );
	GRANT ALL PRIVILEGES ON DATABASE $DB_NAME TO $DB_USER;
    GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO $DB_USER;
EOSQL