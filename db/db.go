package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/luna-duclos/instrumentedsql"
	"log"
	"onsite/utils"
	"strings"
)

func setupLoggerFn() instrumentedsql.LoggerFunc {
	//https://github.com/luna-duclos/instrumentedsql/blob/master/sql_example_test.go
	loggerFn := instrumentedsql.LoggerFunc(func(ctx context.Context, msg string, keyvals ...interface{}) {
		a := msg == "sql-stmt-query" //select
		b := msg == "sql-stmt-exec"  //insert
		if a || b {
			//log.Printf("%s %v", msg, keyvals)
			//4 => duration string, 5=> duration number, 6,7 are args string and value
			//log.Printf("%v %v", keyvals[1], keyvals[7]) //raw sql
			c := keyvals[7].(string)
			//before: {[string "VALUE"], [int 33]} or {[int 33]}
			c = c[1 : len(c)-1]
			//regardless single arg or args
			arrStr := strings.Split(c, ",")
			arr := make([]string, 0)
			for _, v := range arrStr {
				arr = append(arr, utils.UnmarshalSingleTypeValue(v))
			}
			//log.Println("Args: ", arr)

			originSql := keyvals[1].(string)
			//multiple assignments
			for i, times := 0, strings.Count(originSql, "?"); i < times; i++ {
				originSql = strings.Replace(originSql, "?", arr[i], 1)
				//log.Println(originSql)
			}
			log.Println(originSql)
			//log.Printf("%T\n",c)
		}
	})
	return loggerFn
}

const (
	host     = "localhost"
	db       = "db2"
	user     = "tester"
	password = "tester"
)

func New() *sql.DB {
	//wrap driver and set logger
	sql.Register("instrumented-mysql", instrumentedsql.WrapDriver(&mysql.MySQLDriver{}, instrumentedsql.WithLogger(setupLoggerFn())))

	dsn := fmt.Sprintf("%s:%s@%s/%s", user, password, host, db)
	db, err := sql.Open("instrumented-mysql", dsn)
	utils.DieIf(err)

	//boil.DebugMode = true
	return db
}
