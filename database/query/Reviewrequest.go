package query

import (
	"database/sql"
	mydb "review-rest/database/init"
	"fmt"
)

//InfoDeployPool 返回切片类型的数据
//var InfoDeployPool = make([]structerr.DeploymentsPool, 0)

//var ListDeployPool = make([]structerr.DeploymentsPool, 0)

//GetDeployPoolList 获取集群信息列表
func GetDeployPoolList(commit string) []string {

	//var a structerr.DeploymentsPool
	db := mydb.InitDB()
	defer db.Close()
	rows, err := db.Query("SELECT ship_it FROM reviews_review  where review_request_id  = (select id from reviews_reviewrequest  where commit_id = ? ) order   by   id   desc   limit  1 ", commit)
	if err != nil {
		fmt.Println("err")
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}
	//fmt.Println("columns:", columns)
	values := make([]sql.RawBytes, len(columns))
	//fmt.Println("values:", values)
	scanArgs := make([]interface{}, len(values))
	//fmt.Println("scanArgs:", scanArgs)
	for i := range values {
		scanArgs[i] = &values[i]
		//fmt.Println("scanArgs:", scanArgs)
		//fmt.Println("values:", values)
	}
	//fmt.Println("scanArgs:", scanArgs)
	var value []string
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			fmt.Println("log:", err)
			panic(err.Error())
		}
		for _, col := range values {
			//fmt.Println("col:", col)
			if col == nil {
				value = append(value, "NULL")
			} else {
				value = append(value, string(col))
			}
		}
	}
	//fmt.Println("scanArgs:", scanArgs)
	return value
}
