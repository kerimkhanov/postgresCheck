package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

func main() {
	// Чтение переменных среды для учетных данных PostgreSQL.
	host := "main_postgres"
	port := "5432"
	user := "postgresUser"
	password := "pgpwd4postgres"
	dbname := "postgresdb"

	// Формирование строки подключения на основе переменных среды.
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		password,
		host,
		port,
		dbname)

	// Открытие соединения с базой данных.
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("CAN'T Connected")
		return
	}
	defer db.Close()

	// Проверка успешности подключения.
	err = db.Ping()
	if err != nil {

		fmt.Println("CAN'T PING")
		time.Sleep(1 * time.Hour)
		return
	}
	fmt.Println("Успешное подключение к базе данных!")

	// Теперь вы можете выполнять операции с базой данных с помощью 'db'.
	// Например, вы можете выполнять запросы данных:
	rows, err := db.Query("SELECT * FROM your_table")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Итерация по строкам.
	for rows.Next() {
		var id int
		var name string
		// Считывание значений из строки в переменные.
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		// Что-то делаем с данными.
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}
	// Проверка ошибок при итерации по строкам.
	if err = rows.Err(); err != nil {
		panic(err)
	}
}
