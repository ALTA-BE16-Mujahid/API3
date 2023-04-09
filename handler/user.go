package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mujahxd/api3-jwt/auth"
	"github.com/mujahxd/api3-jwt/helper"
	"github.com/mujahxd/api3-jwt/user"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c echo.Context) error {
	var input user.RegisterUserInput

	err := c.Bind(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := echo.Map{"errors": errors}
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	token, err := h.authService.GenerateToken(int(newUser.ID))
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	formatter := user.FormatUser(newUser, token)
	response := helper.APIResponse("Account has been registered", http.StatusCreated, "success", formatter)

	return c.JSON(http.StatusCreated, response)
}

func (h *userHandler) Login(c echo.Context) error {
	var input user.LoginInput

	err := c.Bind(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := echo.Map{"errors": errors}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		log.Error(err)
		errorMessage := echo.Map{"errors": err.Error()}
		response := helper.APIResponse("Login failed", http.StatusBadRequest, "error", errorMessage)
		return c.JSON(http.StatusBadRequest, response)
	}

	token, err := h.authService.GenerateToken(int(loggedinUser.ID))
	if err != nil {
		response := helper.APIResponse("Login account failed", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	formatter := user.FormatUser(loggedinUser, token)
	response := helper.APIResponse("Successfully loggedin", http.StatusOK, "success", formatter)
	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckPhoneAvailability(c echo.Context) error {
	var input user.CheckPhoneInput
	err := c.Bind(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := echo.Map{"errors": errors}
		response := helper.APIResponse("Phone checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	isPhoneAvailable, err := h.userService.IsPhoneAvailable(input)
	if err != nil {
		errorMessage := echo.Map{"errors": "server error"}
		response := helper.APIResponse("Phone checking failed", http.StatusBadRequest, "error", errorMessage)
		return c.JSON(http.StatusBadRequest, response)
	}
	data := echo.Map{
		"is_available": isPhoneAvailable,
	}

	metaMessage := "Phone has been registered"

	if isPhoneAvailable {
		metaMessage = "Phone is available"
	}
	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) UpdateUser(c echo.Context) error {
	inputID, _ := strconv.Atoi(c.Param("id"))

	var inputData user.RegisterUserInput
	err := c.Bind(&inputData)
	if err != nil {
		log.Errorf("kesalahan 1: %s ", err)
		errors := helper.FormatError(err)
		errorMessage := echo.Map{"errors": errors}
		response := helper.APIResponse("Failed to update user", http.StatusUnprocessableEntity, "error", errorMessage)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	updatedUser, err := h.userService.UpdateUser(inputID, inputData)
	if err != nil {
		log.Errorf("kesalahan 2: %s", err)
		response := helper.APIResponse("Failed to update user", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	response := helper.APIResponse("Success to update user", http.StatusOK, "success", updatedUser)
	return c.JSON(http.StatusOK, response)

}
