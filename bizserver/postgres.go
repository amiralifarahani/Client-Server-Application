package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "fmt"

	gen "main/gen/go"

	_ "github.com/lib/pq"
)

func get_user_by_id(input_id string) []*gen.User {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()
	users := make([]*gen.User, 1)
	rows, err := db.Query(`SELECT * FROM "biz_table" WHERE id = $1`, input_id)
	CheckError(err)

	for rows.Next() {
		var name string
		var family string
		var id int32
		var age int32
		var sex string
		var createdAt int64
		err = rows.Scan(&name, &family, &sex, &age, &createdAt, &id)
		//fmt.Println(name, family, age)
		newUser := gen.User{Name: name, Family: family, Age: age, Sex: sex, CreatedAt: createdAt, Id: id}
		//fmt.Println(newUser)
		users[0] = &newUser
	}
	defer rows.Close()
	return users
}

func get_top_100_users() []*gen.User {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	insertStmt := `-- insert into "biz_table" (name, family,age, sex,createdat, id) values ('negar','javadian', 12, 'female', 7800, 14)`
	_, e := db.Exec(insertStmt)
	CheckError(e)

	// close database
	users := make([]*gen.User, 100)
	rows, err := db.Query(`SELECT * FROM "biz_table"`)
	CheckError(err)
	i := 0
	for rows.Next() {
		var name string
		var family string
		var id int32
		var age int32
		var sex string
		var createdAt int64
		err = rows.Scan(&name, &family, &sex, &age, &createdAt, &id)
		//fmt.Println(name, family, age)
		newUser := gen.User{Name: name, Family: family, Age: age, Sex: sex, CreatedAt: createdAt, Id: id}
		//fmt.Println(newUser)
		users[i] = &newUser
		i++
	}
	defer rows.Close()
	defer db.Close()
	return users

	//fmt.Println(name, family, age)

	//CheckError(err)
}
