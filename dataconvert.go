package dataconvert

var SQL string

func AccessMysql(datasoucename string) {
	amysql := accessMysqlDB{dataSourceName: datasoucename}
	amysql.Open()
	defer amysql.Close()
	SQL = "select id from conversation where id = 1"
	amysql.FetchData(SQL)
}
