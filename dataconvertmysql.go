package dataconvert

import (
	"database/sql"
	"fmt"
)

var mysqlDB *sql.DB

type accessMysqlDB struct {
	dataSourceName string
}

func (access accessMysqlDB) Open() {
	var err error
	//Open database connection
	mysqlDB, err = sql.Open("mysql", access.dataSourceName)
	if err != nil {
		panic(err.Error())
	}
}

func (accessMysqlDB) Close() {
	//Close database connection
	mysqlDB.Close()
}

func (accessMysqlDB) FetchData(SQL string) {
	//query data from mysql
	datarows, err := mysqlDB.Query(SQL)
	if err != nil {
		panic(err.Error())
	}

	//Get column names
	columns, err := datarows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]string, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for datarows.Next() {
		err := datarows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string
		for _, col := range values {
			if col == "" {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(col, ":", value)
		}
		fmt.Println("------------------------")
	}
	if err = datarows.Err(); err != nil {
		panic(err.Error())
	}
}
