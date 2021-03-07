package main

import (
	"fmt"
	"net/http"

	"github.com/nikydobrev/go-rest-api-demo/internal/database"
	transportHTTP "github.com/nikydobrev/go-rest-api-demo/internal/transport/http"
)

// App - the struct which contains things like
// pointers to database connections
type App struct{}

// Run - handles the startup of our application
func (a *App) Run() error {
	fmt.Println("Setting Up Our App")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	fmt.Println(db.DB())

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()
	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up")
		fmt.Println(err)
	}
}
