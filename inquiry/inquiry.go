package inquiry

import (
	"database/sql"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

// 問い合わせ項目
type InquiryItem struct {
	ID     string `json:"id", primarykey, autoincrement`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

// list
func List() []InquiryItem {
	dbmap := open()
	defer dbmap.Db.Close()

	var items []InquiryItem
	_, err := dbmap.Select(&items, "select * from inquiry")
	if err != nil {
		panic(err.Error())
	}

	return items
}

// save
func (inquiry InquiryItem) Save() {
	dbmap := open()
	defer dbmap.Db.Close()

	err := dbmap.Insert(inquiry)
	if err != nil {
		panic(err.Error())
	}
}

func open() *gorp.DbMap {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/inquiry?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	return dbmap
}
