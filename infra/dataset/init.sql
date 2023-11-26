CREATE TABLE inflation_data (
    ID INT,
    Name TEXT,
    AGE INT,
);

COPY inflation_data
FROM '/docker-entrypoint-initdb.d/table_1.csv'
DELIMITER ','
CSV HEADER;