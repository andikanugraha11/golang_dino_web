package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type animal struct {
	id         int
	animalType string
	name       string
	zone       int
	age        int
}

func main() {
	db, err := sql.Open("mysql", "root:detikcom@/learn_golang")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Query with argument
	rows, err := db.Query("select * from learn_golang.animal")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	animals := []animal{}
	for rows.Next() {
		a := animal{}
		err := rows.Scan(&a.id, &a.animalType, &a.name, &a.zone, &a.age)
		if err != nil {
			log.Println(err)
			continue
		}

		animals = append(animals, a)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(animals)

}
