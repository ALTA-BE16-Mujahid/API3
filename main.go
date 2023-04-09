package main

import (
	"log"
	"net/http"
	"strings"

	// "github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/mujahxd/api3-jwt/auth"
	book "github.com/mujahxd/api3-jwt/book"
	"github.com/mujahxd/api3-jwt/config"
	"github.com/mujahxd/api3-jwt/handler"
	"github.com/mujahxd/api3-jwt/helper"
	user "github.com/mujahxd/api3-jwt/user"
)

func main() {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	// database
	db := config.ConnectionDB(&loadConfig)
	// validate := validator.New()

	db.AutoMigrate(&user.User{}, &book.Book{})

	userRepository := user.NewRepository(db)
	bookRepository := book.NewRepository(db)

	userService := user.NewService(userRepository)
	bookService := book.NewService(bookRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	bookHandler := handler.NewBookHandler(bookService)

	e := echo.New()
	api := e.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/phone_checkers", userHandler.CheckPhoneAvailability)
	api.PUT("/users/:id", userHandler.UpdateUser, authMiddleware(authService, userService))

	api.GET("/books", bookHandler.GetBooks)
	api.POST("/books", bookHandler.CreateBook, authMiddleware(authService, userService))
	api.PUT("/books/:id", bookHandler.UpdateBook, authMiddleware(authService, userService))
	api.DELETE("/books/:id", bookHandler.DeleteBook, authMiddleware(authService, userService))

	e.Start(":8080")

}

func authMiddleware(authService auth.Service, userService user.Service) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")

			if !strings.Contains(authHeader, "Bearer") {
				response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
				return c.JSON(http.StatusUnauthorized, response)
			}

			// bearer {token}
			var tokenString string
			arrayToken := strings.Split(authHeader, " ")
			if len(arrayToken) == 2 {
				tokenString = arrayToken[1]
			}

			token, err := authService.ValidateToken(tokenString)
			if err != nil {
				response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
				return c.JSON(http.StatusUnauthorized, response)
			}

			claim, ok := token.Claims.(jwt.MapClaims)

			if !ok || !token.Valid {
				response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
				return c.JSON(http.StatusUnauthorized, response)
			}

			userID := int(claim["user_id"].(float64))
			user, err := userService.GetUserByID(userID)
			if err != nil {
				response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
				return c.JSON(http.StatusUnauthorized, response)
			}

			c.Set("currentUser", user)
			return next(c)
		}
	}
}

// func authMiddleware1(authService auth.Service, userService user.Service) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		authHeader := c.Request().Header.Get("Authorization")

// 		if !strings.Contains(authHeader, "Bearer") {
// 			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
// 			return c.JSON(http.StatusUnauthorized, response)
// 		}

// 		// bearer {token}
// 		var tokenString string
// 		arrayToken := strings.Split(authHeader, " ")
// 		if len(arrayToken) == 2 {
// 			tokenString = arrayToken[1]
// 		}

// 		token, err := authService.ValidateToken(tokenString)
// 		if err != nil {
// 			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
// 			return c.JSON(http.StatusUnauthorized, response)
// 		}

// 		claim, ok := token.Claims.(jwt.MapClaims)

// 		if !ok || !token.Valid {
// 			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
// 			return c.JSON(http.StatusUnauthorized, response)
// 		}

// 		userID := int(claim["user_id"].(float64))
// 		user, err := userService.GetUserByID(userID)
// 		if err != nil {
// 			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
// 			return c.JSON(http.StatusUnauthorized, response)
// 		}

// 		c.Set("currentUser", user)
// 		return nil

// 	}

// }
