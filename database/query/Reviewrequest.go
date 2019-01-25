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

func GetCommitId(commit string) []string {

        //var a structerr.DeploymentsPool
        db := mydb.InitDB()
        defer db.Close()
        rows, err := db.Query("SELECT commit_id FROM reviews_reviewrequest where commit_id = ? ", commit)
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
func GetCommitReviewId(commit string) []string {

        //var a structerr.DeploymentsPool
        db := mydb.InitDB()
        defer db.Close()
        rows, err := db.Query("SELECT id FROM reviews_reviewrequest where commit_id = ? ", commit)
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
func GetReviewId() []string {

	db := mydb.InitDB()
	defer db.Close()
	rows, err := db.Query("SELECT id FROM reviews_reviewrequest WHERE public = '1' AND status = 'p' AND time_added  >= DATE_SUB(CURDATE(), INTERVAL 7 DAY)")
	if err != nil {
		fmt.Println("err")
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))

	for i := range values {
		scanArgs[i] = &values[i]
	}
	var value []string
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			fmt.Println("log:", err)
			panic(err.Error())
		}
		for _, col := range values {
			if col == nil {
				value = append(value, "NULL")
			} else {
				value = append(value, string(col))
			}
		}
	}
	return value
}

func PostSummary(summary string, reviewid string) []string {

        //var a structerr.DeploymentsPool
        db := mydb.InitDB()
        defer db.Close()
        rows, err := db.Query("update reviews_reviewrequest set summary = ? where id = ?", summary, reviewid)
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
