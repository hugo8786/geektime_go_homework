package dao

import (
	"fmt"
	"time"

	"hugo.com/geektime/error/db"
)

var dbConn *db.DB

const (
	testMysqlHost    = "127.0.0.1"
	testMysqlPort    = "3306"
	testMysqlAccount = "root"
	testMysqlPwd     = "Mslnu#@QNxgIM"
	testMyDbName     = "test_dao"
)

func init() {
	db, err := db.ConnectSQL(testMysqlHost, testMysqlPort, testMysqlAccount, testMysqlPwd, testMyDbName)
	if err != nil {
		panic(fmt.Sprintf("db connect faild , err= %#v", err))
	} else {
		fmt.Printf("db %s connect success", testMyDbName)
		dbConn = db
	}
}

type TestDaoResult struct {
	Id         int
	Name       string
	CreateTime time.Time
}

func TestDao() (db *db.DB) {
	sql := "select id ,name,create_time from dao_result "
	db = dbConn.Clone()
	rows, err := db.Sql.Query(sql)
	dbConn.AddError(err)
	defer rows.Close()
	results := make([]*TestDaoResult, 0)
	db.Value = results
	db.RowsAffected = 0
	for rows.Next() {
		result := &TestDaoResult{}
		db.RowsAffected++
		db.AddError(rows.Scan(
			result.Id,
			result.Name,
			result.CreateTime,
		))
		results = append(results, result)
	}
	return
}
func main() {

	db := TestDao()
	if db.Error != nil {
		fmt.Printf("query error = %#v", db.Error)
		return
	}
	if db.RowsAffected == 0 {
		fmt.Printf("query result is Empty ")
		return
	}
	fmt.Printf("query result = %#v", db.Value)
}
