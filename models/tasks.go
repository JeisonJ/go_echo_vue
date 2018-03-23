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
	Name string `json:"name"`
}

// TaskCollection es una colección de tareas.
type TaskCollection struct {
	Task []Task `json:"items"`
}

func GetTasks(db *sql.DB) TaskCollection {
	sql := "SELECT * FROM tasks"

	rows, err := db.Query(sql)

	// Salir si la sentencia SQL no se ejecuta por alguna razón.
	if err != nil {
		panic(err)
	}

	// Cerrar la consulta al finalizar todo.
	defer rows.Close()

	// Resultados
	result := TaskCollection{}
	for rows.Next()  {
		task := Task{}
		err2 := rows.Scan(&task.ID, &task.Name)

		// Si existe algún error detener la ejecución
		if err2 != nil {
			panic(err2)
		}
		result.Task = append(result.Task, task)
	}
	return result
}

// PutTask inserta una nueva tarea en la base de datos y retorna la nueva
// id si toso ha ido bien, en caso contrario detiene la ejecución.
func PutTask(db *sql.DB, name string) (int64, error) {
	sql := "INSERT INTO tasks(name) VALUES(?)"

	// Crear una sentencia SQL preparada.
	stmt, err := db.Prepare(sql)

	// Salir si se obtiene un error.
	if err != nil {
		panic(err)
	}

	// Cerrar la consulta al finalizar todo.
	defer stmt.Close()

	// Remplazar (?) con los valores
	result, err2 := stmt.Exec(name)

	// Salir si hay error
	if err2 != nil {
		panic(err2)
	}
	// Retornar el ultimo id insertado/generado.
	return result.LastInsertId()
}


func DeleteTask(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM tasks WHERE id = ?"

	// Crear una sentencia SQL preparada.
	stmt, err := db.Prepare(sql)
	// Salir si se obtiene un error.
	if err != nil {
		panic(err)
	}

	// Remplazar (?) con los valores
	result, err2 := stmt.Exec(id)
	// Salir si hay error
	if err2 != nil {
		panic(err2)
	}
	// Retornar la fila afectada.
	return result.RowsAffected()
}