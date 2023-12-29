
-- drop if exists
DROP DATABASE IF EXISTS infinity_db WITH (FORCE);
DROP ROLE IF EXISTS test_user;

-- create users and db
DROP SCHEMA IF EXISTS infinity ;
CREATE USER test_user WITH PASSWORD 'password';
CREATE DATABASE infinity_db;
GRANT ALL PRIVILEGES ON DATABASE "infinity_db" to test_user;
