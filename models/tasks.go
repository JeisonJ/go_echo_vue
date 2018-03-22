package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

/*
 En lugar de simplemente hacer llamadas directas a la base de
 datos de nuestros controladores, mantendremos
 nuestro código ordenado y ordenado al abstraer la lógica
 de la base de datos en un modelo.
*/

// Task es una estructura que contiene datos de la Tarea.
type Task struct {
	ID   int    `json:"id"`
	name string `json:"name"`
}

// TaskCollection es una colección de tareas.
type TaskCollection struct {
	Task []Task `json:"items"`
}

func GetTask(db *sql.DB) TaskCollection {

}

func PutTask(db *sql.DB, name string) (int64, error) {

}
