package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Db struct {
	Host     string
	User     string
	Password string
	Database string
	Port     string
	Conn     *gorm.DB
}

func (db *Db) NewConnection() {
	dns := db.User + ":" + db.Password + "@tcp(" + db.Host + ":" + db.Port + ")/" + db.Database
	fmt.Println(dns)
	conn, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dns,
	}))

	if err != nil {
		PrintDbInfo("Error connecting to database")
		return
	}

	db.Conn = conn
}

func PrintDbInfo(info string) {
	fmt.Println("[red-social-db] " + info)
}
