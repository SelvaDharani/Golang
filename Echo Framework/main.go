package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.office.opendns.com/selthiru/Golang-Training/db"
	"github.office.opendns.com/selthiru/Golang-Training/handlers"
	jwtmiddleware "github.office.opendns.com/selthiru/Golang-Training/middleware"
	"github.office.opendns.com/selthiru/Golang-Training/view"
)

// Custom middleware to add a header to all responses
func customMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("X-Custom-Middleware", "Hello from custom middleware!")
		return next(c)
	}
}

func main() {
	// Entrypoint - Create a new Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(customMiddleware)

	// Initialize the database connection
	mysqlBackend := db.MysqlBackend{}
	dbInstance, err := mysqlBackend.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize the database:", err)
	}
	// Custom template renderer
	renderer := view.NewTemplateRenderer()
	e.Renderer = renderer

	// Routes
	e.GET("/", func(c echo.Context) error {
		// Fetch all users from the database
		users, err := handlers.GetAllUsersHandler(c, dbInstance)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to fetch users")
		}

		// Render the index.html template with the list of users
		return c.Render(http.StatusOK, "index.html", users)
	})

	//Routes
	e.GET("/health", handlers.IndexHandler)
	//e.GET("/users", func(c echo.Context) error {
	//	return handlers.GetAllUsersHandler(c, dbInstance) // Pass the database connection to the handler
	//})
	//e.POST("/users", func(c echo.Context) error {
	//		return handlers.CreateUserHandler(c, dbInstance) // Pass the database connection to the handler
	//})
	e.POST("/users", func(c echo.Context) error {
		return handlers.CreateUserHandler(c, dbInstance)
	}, jwtmiddleware.JWTMiddleware)
	e.PUT("/users", func(c echo.Context) error {
		return handlers.UpdateUserHandler(c, dbInstance) // Pass the database connection to the handler
	})
	e.DELETE("/users", func(c echo.Context) error {
		return handlers.DeleteUserHandler(c, dbInstance) // Pass the database connection to the handler
	})
	// Entrypoint -  Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
