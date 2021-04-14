package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Db struct {
	host     string
	user     string
	password string
	database string
	port     string
	conn     *gorm.DB
}

func (db *Db) NewConnection() {
	dns := db.user + ":" + db.password + "@tcp(" + db.host + ":" + db.port + ")/" + db.database
	fmt.Println(dns)
	conn, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dns,
	}))

	if err != nil {
		PrintDbInfo("Error connecting to database")
		return
	}

	db.conn = conn
}

func PrintDbInfo(info string) {
	fmt.Println("[red-social-db] " + info)
}
