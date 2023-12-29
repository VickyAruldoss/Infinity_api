#!usr/bin/env bash

echo "Running liquibase Migrations for infinity"
pushd ./database || exit

DB_URL="jdbc:postgresql://localhost:5432/infinity_db?user=test_user&password=password"
DB_CHANGE_LOG_FILE="changelog.xml"
DB_SCHEMA="infinity"

liquibase --url="${DB_URL}" --contexts="dev" --liquibase-schema-name="${DB_SCHEMA}" --changelog-file="${DB_CHANGE_LOG_FILE}" update

popd ..
echo "complete migrations"