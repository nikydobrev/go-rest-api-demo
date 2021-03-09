package main

import (
	"fmt"
	"net/http"

	"github.com/nikydobrev/go-rest-api-demo/internal/comment"
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
	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
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
