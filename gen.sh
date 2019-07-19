#!/usr/bin/env bash

echo "update toml file"
sqlboiler --version

#rm models/*

#sqlboiler mssql
#sqlboiler --wipe --no-hooks --no-context --debug mssql
sqlboiler --wipe --no-hooks --no-context mssql

echo "testing"
#cp tables_schema.sql models/tables_schema.sql
go test ./models

