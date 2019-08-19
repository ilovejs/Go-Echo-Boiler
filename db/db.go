package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/luna-duclos/instrumentedsql"
	"github.com/volatiletech/sqlboiler/boil"
	"log"
	"onsite/utils"
	"strings"
)

func setupLoggerFn() instrumentedsql.LoggerFunc {
	logFn := instrumentedsql.LoggerFunc(
		func(ctx context.Context, msg string, keyvals ...interface{}) {
			// fmt.Println("Driver Logger hits, otherwise driver issue")

			a := msg == "sql-stmt-query" // select
			b := msg == "sql-stmt-exec"  // insert
			if a || b {
				// log.Printf("%s %v", msg, keyvals)
				// 4 => duration string, 5=> duration number, 6,7 are args string and value
				// log.Printf("%v %v", keyvals[1], keyvals[7]) //raw sql
				c := keyvals[7].(string)
				// before: {[string "VALUE"], [int 33]} or {[int 33]}
				c = c[1 : len(c)-1]
				// regardless single arg or args
				arrStr := strings.Split(c, ",")
				arr := make([]string, 0)
				for _, v := range arrStr {
					arr = append(arr, utils.UnmarshalSingleTypeValue(v))
				}
				// log.Println("Args: ", arr)

				originSql := keyvals[1].(string)
				// multiple assignments, mysql style ?, mssql style $1
				for i, times := 0, strings.Count(originSql, "$"); i < times; i++ {
					originSql = strings.Replace(originSql, "$", arr[i], 1)
					// log.Println(originSql)
				}
				log.Println(originSql)
				// log.Printf("%T\n",c)
			}
		})
	// logFn.Log(ctx, msg, keyvals)
	return logFn
}

// todo: env
const (
	db = "test2"
	// host     = "localhost"
	// host     = "gateway.docker.internal"
	host     = "host.docker.internal"
	port     = "1433"
	user     = "tester"
	password = "tester"
	// driverName = "instrumented-mssql"
	driverName = "mssql"
)

func New() *sql.DB {
	// logger := setupLoggerFn()

	// wrap driver with logger
	// sql.Register(driverName,
	//	instrumentedsql.WrapDriver(&mssql.Driver{}, //check
	//		instrumentedsql.WithLogger(logger)))

	// sqlserver://username:password@host/instance?param1=value&param2=value
	// dsn := "sqlserver://tester:tester@localhost:1433?database=test2"

	// TODO: connection timeout 30
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", user, password, host, port, db)
	log.Println(dsn)

	db, err := sql.Open(driverName, dsn)
	utils.DieIf(err)
	log.Println("Database is Connected. Check driver if any issue.")

	// default sql debugger
	boil.DebugMode = true
	return db
}
