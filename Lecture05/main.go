package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	dataForDb := "user=mrtsk dbname=exampledb sslmode=disable password=verystrongpassword" // данные для подключения к БД
	db, err := sql.Open("postgres", dataForDb)                                             // подключаемся к БД
	if err != nil {                                                                        // если ошибка, то выводим ее и завершаем работу
		log.Fatal(err)
	}
	defer db.Close()

	// Создаем таблицу
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS people (id SERIAL PRIMARY KEY, name VARCHAR(50))")
	if err != nil {
		log.Fatal(err)
	}

	// Вставляем данные
	_, err = db.Exec("INSERT INTO people (name) VALUES ($1)", "Marat Askarbek")
	if err != nil {
		log.Fatal(err)
	}

	// Получаем данные
	rows, err := db.Query("SELECT id, name FROM people")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d: %s\n", id, name)
	}

	// Проверяем наличие ошибок после Next или rows.Close
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}
