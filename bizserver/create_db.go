package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "postgres_db"
	port     = 5432
	user     = "postgres"
	password = "12345"
	dbname   = "biz_database"
)

// func CreateServer() *fiber.App {
// 	app := fiber.New()

// 	return app
// }

// func main() {

// 	//app := CreateServer()
// 	//
// 	//app.Use(cors.New())
// 	//
// 	//app.Get("/hello", func(c *fiber.Ctx) error {
// 	//	create_database()
// 	//	return c.SendString("Hello, World!")
// 	//})
// 	//
// 	//// 404 Handler
// 	//app.Use(func(c *fiber.Ctx) error {
// 	//	return c.SendStatus(404) // => 404 "Not Found"
// 	//})
// 	//
// 	//log.Fatal(app.Listen(":3000"))

// 	Create_database()
// }

func Create_database() {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	// if err != nil{
	// 	fmt.Println(err)
	// }
	CheckError(err)
	_, err = db.Query(`CREATE TABLE biz_table (name TEXT, family TEXT,sex TEXT,age bigint,createdat bigint,id bigint);`)
	CheckError(err)
	insertStmt := `-- insert into biz_table (name, family,age, sex,createdat, id) values ('par','javadian', 12, 'female', 7800, 1)`
	_, err = db.Exec(insertStmt)
	CheckError(err)
	insertStmt = `-- insert into biz_table (name, family,age, sex,createdat, id) values ('negar','bsh', 12, 'female', 7800, 2)`
	_, err = db.Exec(insertStmt)
	CheckError(err)
	insertStmt = `-- insert into biz_table (name, family,age, sex,createdat, id) values ('kiana','msz', 12, 'female', 7800, 3)`
	_, err = db.Exec(insertStmt)
	CheckError(err)
	//dbName := "biz_database3"
	//_, err = db.Exec("create database " + dbname)
	//CheckError(err)

	// psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	//// open database
	//db, err = sql.Open("postgres", psqlconn)
	//CheckError2(err)

	//todo :code below creates a table
	// _, err = db.Query(`CREATE TABLE "biz_table" (name TEXT, family TEXT,sex TEXT,age bigint,createdat bigint,id bigint);`)
	// CheckError2(err)

	// insertStmt = `-- insert into "biz_table" (name, family,age, sex,createdat, id) values ('negar','javadian', 12, 'female', 7800, 14)`
	// _, e = db.Exec(insertStmt)
	// CheckError2(e)
	defer db.Close()


}

func CheckError2(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
