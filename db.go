package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//fmt.Println("hello word!")
	query()
}

func insert() {

}

//123.57.16.143

func query() {
	db, err := sql.Open("mysql", "root:hai0825@tcp(47.104.28.223:3306)/test?charset=utf8")
	checkErr(err)

	rows, err := db.Query("SELECT id,username FROM user")
	checkErr(err)

	//普通
	// for rows.Next() {
	// 	var id int
	// 	var username string

	// 	rows.Columns()
	// 	err = rows.Scan(&id, &username)
	// 	checkErr(err)

	// 	fmt.Println(id, username)

	// }

	//字典类型
	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
	//fmt.Println(record)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
