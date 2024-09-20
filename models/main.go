package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	id             int
	branch_code    string
	branch_name    string
	pickup_code    string
	courier_code   string
	courier_name   string
	courier_status string
	bpc_code       string
	create_time    uint32
	update_time    uint32
}

// 查询数据一般使用 Scan
// 缺点，不通用，不同的表的查询需要定义不同的结构体，然后写不同的 Scan。
func searchBook1(db *sql.DB) ([]User, error) {
	var userss []User
	//查询语句
	sql := `select id,branch_code,branch_name from shipment_code_courier`
	rows, err := db.Query(sql)
	if err != nil {
		log.Panicln(err)
		return userss, err
	}
	//defer rows.Close()
	for rows.Next() {
		var user1 User
		// 获取各列的值，放到对应的地址中
		rows.Scan(&user1.id, &user1.branch_code, &user1.branch_name)
		userss = append(userss, user1)
	}
	return userss, nil
}

func main() {
	//Go连接Mysql
	//用户名:密码啊@tcp(ip:端口)/数据库的名字
	dsn := "root:123456@tcp(127.0.0.1:3306)/st"
	//连接数据集
	db, err := sql.Open("mysql", dsn) //open不会检验用户名和密码
	if err != nil {
		fmt.Printf("dsn:%s invalid,err:%v\n", dsn, err)

	}
	err = db.Ping() //尝试连接数据库
	if err != nil {
		fmt.Printf("open %s faild,err:%v\n", dsn, err)
	}
	fmt.Println("连接数据库成功~")
	if err != nil {
		fmt.Printf("init DB failed,err%v\n", err)
	}
	book1, err := searchBook1(db)
	fmt.Print(book1, err)

}
