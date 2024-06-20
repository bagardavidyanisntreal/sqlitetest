package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	os.Remove("sqlite-database.db") // I delete the file to avoid duplicated records.

	log.Println("Creating sqlite-database.db...")
	file, err := os.Create("sqlite-database.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db created")

	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db")
	defer sqliteDatabase.Close()

	createTable(sqliteDatabase)

	// INSERT RECORDS
	insertTask(sqliteDatabase, ptr("Liana Kim"))
	insertTask(sqliteDatabase, nil)
	insertTask(sqliteDatabase, nil)
	insertTask(sqliteDatabase, ptr("Alayna Armitage"))
	insertTask(sqliteDatabase, ptr("Marni Benson"))
	insertTask(sqliteDatabase, nil)
	insertTask(sqliteDatabase, ptr("Leigh Daly"))
	insertTask(sqliteDatabase, ptr("Marni Benson"))
	insertTask(sqliteDatabase, ptr("Klay Correa"))

	// DISPLAY INSERTED RECORDS
	displayStudents(sqliteDatabase)
}

func ptr[T any](data T) *T {
	return &data
}

func createTable(db *sql.DB) {
	createStudentTableSQL := `CREATE TABLE task (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT
	  );`

	statement, err := db.Prepare(createStudentTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("student table created")
}

func insertTask(db *sql.DB, name *string) {
	insertStudentSQL := `INSERT INTO task(name) VALUES (?)`
	statement, err := db.Prepare(insertStudentSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = statement.Exec(name)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func displayStudents(db *sql.DB) {
	row, err := db.Query("SELECT * FROM task")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	var tasks []Task
	for row.Next() {
		t := Task{}
		err = row.Scan(&t.ID, &t.Name)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		tasks = append(tasks, t)
	}

	for _, task := range tasks {
		fmt.Println(task)
	}
}

type Task struct {
	ID   int
	Name *string
}
