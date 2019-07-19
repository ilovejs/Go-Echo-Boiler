#!/usr/bin/env bash

echo "update toml file"
sqlboiler --version

rm models/*

#sqlboiler mssql
#sqlboiler --wipe --no-hooks --no-context --debug mssql
sqlboiler --wipe --no-hooks --no-context mssql

cp tables_schema.sql models/tables_schema.sql
echo "tables_schema.sql copied"

cp mssql_main_test.go models/mssql_main_test.go
echo "patch generated code"

echo "testing"
#go test ./models

