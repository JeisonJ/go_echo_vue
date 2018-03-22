package main

import (
	// Paquete estandar http
	"net/http"

	// Paquete estandar para manejo de base de datos
	"database/sql"

	// Funciones para resolver las rutas.
	"github.com/jeisonj/go-echo-vue/handlers"

	// Driver para sqlite3
	_"github.com/mattn/go-sqlite3"

	// Echo framework
	"github.com/labstack/echo"
	_"github.com/labstack/echo/middleware"
)

func main() {

	// iniciaremos la base de datos y especificaremos un nombre de archivo
	// de "storage.db". Si este archivo aún no existe, el controlador continuará
	// y lo creará para nosotros.
	database := initDB("storage.db")
	migrate(database)

	// Nueva instancia de Echo.
	e := echo.New()

	// Rutas.
	e.File("/", 	   "public/index.html")
	e.GET("/tasks", 	   handlers.GetTasks(database))
	e.PUT("/tasks", 	   handlers.UpdateTask(database))
	e.DELETE("/task/:id", handlers.DeleteTask(database))

	// Iniciar el servidor web.
	e.Logger.Fatal(e.Start(":8080"))
}


// Funciones para el manejo de la base de datos.
func initDB(filepath string) *sql.DB  {
	database, err := sql.Open("sqlite3", filepath)

	// Manejar un posible error.
	if err != nil {
		panic(err)
	}

	// Si no obtenemos ningún error, pero de alguna manera todavía no
	// logramos una conexión con la db salimos también.
	if database == nil {
		panic("database nil")
	}

	return database
}

func migrate(database *sql.DB)  {

	sql := `
		CREATE TABLE IF NOT EXISTS tasks(
			id   INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name VARCHAR NO NULL
		);
	`

	// Ejecutando la sentencia SQL.
	_, err := database.Exec(sql)


	// Si ha ocurrido algun error al intentar ejecutar la sentencia anterior, salir.
	if err != nil {
		panic(err)
	}

}