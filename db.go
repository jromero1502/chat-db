package db

import (
	"fmt"
	"time"

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

func main() {
	conn := &Db{
		host:     "localhost",
		user:     "jromero1502",
		password: "$Root0123",
		port:     "3306",
		database: "red_social",
	}
	conn.NewConnection()
	u := &User{
		Name:      "Julian",
		Lastname:  "Romero",
		Email:     "romerojulian115@gmail.com",
		Birthday:  time.Now(),
		LastLogin: time.Now(),
	}
	conn.conn.AutoMigrate(&User{})

	result := conn.conn.Create(u)
	if result.Error != nil {
		PrintDbInfo("Error creating user :(")
		return
	}

	id := fmt.Sprint(u.ID)
	PrintDbInfo("User created. ID: " + id)
}
