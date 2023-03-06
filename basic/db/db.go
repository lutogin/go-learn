package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type person struct {
	id   int
	name string
	age  int
}

func RunSql() {
	// connect
	db, err := sql.Open("mysql", "root:Password1@tcp4(127.0.0.1:3306)/learn-go")

	if err != nil {
		fmt.Printf("Error during connection to the DB: %s", err)
	}

	defer db.Close()
	// Exec for INSERT, UPDATE, DELETE, or CREATE TABLE. It returns a sql.Result
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS test (`id` serial,`name` varchar(255) NOT NULL,`age` int NOT NULL, PRIMARY KEY (id))")

	if err != nil {
		fmt.Printf("Error during create a TEST table to the DB: %s \n", err)
	} else {
		fmt.Println("Table was created or used!")
	}

	// Query for SELECT, it returns *sql.Rows
	rows, err := db.Query("SELECT * FROM test WHERE age > ? AND name != ?", 20, "Vasilos")

	defer rows.Close()

	persons := []person{}

	for rows.Next() {
		//var id int
		//var name string
		//var age string
		p := person{}
		err := rows.Scan(&p.id, &p.name, &p.age)

		if err != nil {
			fmt.Printf("Error during scan: %s", err)
			continue
		}
		persons = append(persons, p)
	}

	for _, p := range persons {
		fmt.Printf("ID: %d, NAME: %s, AGE: %d \n", p.id, p.name, p.age)
	}

	err = rows.Err()
	if err != nil {
		fmt.Printf("Error during request to the DB: %s", err)
	}
}

type Product struct {
	Id      bson.ObjectId `bson:"_id"`
	Model   string        `bson:"model"`
	Company string        `bson:"company"`
	Price   int           `bson:"price"`
}

func RunMongo() {
	// открываем соединение
	session, err := mgo.Dial("mongodb://user:password1@127.0.0.1:27017/go-learn")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// получаем коллекцию
	productCollection := session.DB("go-learn").C("products")

	p1 := &Product{Id: bson.NewObjectId(), Model: "iPhone 8", Company: "Apple", Price: 64567}
	// добавляем один объект
	err = productCollection.Insert(p1)
	if err != nil {
		fmt.Println(err)
	}

	p2 := &Product{Id: bson.NewObjectId(), Model: "Pixel 2", Company: "Google", Price: 58000}
	p3 := &Product{Id: bson.NewObjectId(), Model: "Xplay7", Company: "Vivo", Price: 49560}
	// добавляем два объекта
	err = productCollection.Insert(p2, p3)
	if err != nil {
		fmt.Println(err)
	}
}
