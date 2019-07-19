#!/usr/bin/env bash


echo "WARNING: Please update toml file for Credentials and Database"
echo "Init database use script in doc/sql/db_side_run.sql"

sqlboiler --version

rm models/*

#sqlboiler mssql
#sqlboiler --wipe --no-hooks --no-context --debug mssql

sqlboiler --wipe --no-hooks --no-context mssql
echo "model has been reset"

cp tables_schema.sql models/tables_schema.sql
echo "tables_schema.sql copied to models/"

cp mssql_main_test.txt models/mssql_main_test.go
echo "patched bugs test code"

echo "testing"
go test ./models