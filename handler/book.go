package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mujahxd/api3-jwt/book"
	"github.com/mujahxd/api3-jwt/helper"
	"github.com/mujahxd/api3-jwt/user"
)

type bookHandler struct {
	service book.Service
}

func NewBookHandler(service book.Service) *bookHandler {
	return &bookHandler{service}
}

// api/v1/books
func (h *bookHandler) GetBooks(c echo.Context) error {
	userID, _ := strconv.Atoi(c.QueryParam("user_id"))

	books, err := h.service.GetBooks(userID)
	if err != nil {
		response := helper.APIResponse("Error to get books", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	response := helper.APIResponse("List of books", http.StatusOK, "success", book.FormatBooks(books))
	return c.JSON(http.StatusOK, response)
}

func (h *bookHandler) CreateBook(c echo.Context) error {
	var input book.CreateBookInput

	err := c.Bind(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := echo.Map{"errors": errors}
		response := helper.APIResponse("Failed to create book", http.StatusUnprocessableEntity, "error", errorMessage)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	currentUser := c.Get("currentUser").(user.User)
	input.User = currentUser

	newBook, err := h.service.CreateBook(input)
	if err != nil {
		response := helper.APIResponse("Failed to create book", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	response := helper.APIResponse("Success to create book", http.StatusCreated, "success", book.FormatBook(newBook))
	return c.JSON(http.StatusOK, response)

}

func (h *bookHandler) UpdateBook(c echo.Context) error {
	inputID, _ := strconv.Atoi(c.Param("id"))

	var inputData book.CreateBookInput
	err := c.Bind(&inputData)
	if err != nil {
		log.Errorf("kesalahan 1: %s ", err)
		errors := helper.FormatError(err)
		errorMessage := echo.Map{"errors": errors}
		response := helper.APIResponse("Failed to update book", http.StatusUnprocessableEntity, "error", errorMessage)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	updatedBook, err := h.service.UpdateBook(inputID, inputData)
	if err != nil {
		log.Errorf("kesalahan 2: %s", err)
		response := helper.APIResponse("Failed to update book", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	response := helper.APIResponse("Success to update book", http.StatusOK, "success", book.FormatBook(updatedBook))
	return c.JSON(http.StatusOK, response)

}
func (h *bookHandler) DeleteBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)
		response := helper.APIResponse("failed to delete book", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	deletedBook, err := h.service.DeleteBook(id)
	if err != nil {
		log.Error(err)
		response := helper.APIResponse("failed to delete book", http.StatusInternalServerError, "error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helper.APIResponse("book has been deleted", http.StatusOK, "success", book.FormatBook(deletedBook))
	return c.JSON(http.StatusOK, response)
}
