package userv1

import (
	"github.com/addixit1/fiber-boilerplate/internal/config"
	"github.com/addixit1/fiber-boilerplate/internal/querybuilder"
	"github.com/addixit1/fiber-boilerplate/internal/utils/errortracker"
	"github.com/gofiber/fiber/v2"
)

// getLang gets language from context
func getLang(c *fiber.Ctx) string {
	lang, ok := c.Locals("lang").(string)
	if !ok || lang == "" {
		return "en" // Default to English
	}
	return lang
}

// ListUsers godoc
// @Summary List users
// @Description Get all users with search
// @Tags Users
// @Accept json
// @Produce json
// @Param search query string false "Search by name"
// @Param platform header string false "Device OS: 1-Android, 2-iOS, 3-WEB" Enums(1,2,3) default(1)
// @Param timezone header string false "Time zone" default(Asia/Kolkata)
// @Param offset header integer false "Time zone offset" default(0)
// @Param accept-language header string false "Language: en, hi" Enums(en,hi) default(en)
// @Param appversion header string false "App version" default(v1)
// @Param routeversion header string false "Route version" default(v1)
// @Security BearerAuth
// @Success 200 {object} config.APIResponse
// @Failure 401 {object} config.APIResponse
// @Router /users [get]
func List(c *fiber.Ctx) error {
	lang := getLang(c)

	filter := querybuilder.New().
		Regex("name", c.Query("search")).
		Build()

	users, err := ListUsers(filter)
	if err != nil {
		return c.Status(500).JSON(config.InternalServerError(lang))
	}

	return c.Status(200).JSON(config.List(users, lang))
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param body body CreateUserDTO true "Create user"
// @Param platform header string false "Device OS: 1-Android, 2-iOS, 3-WEB" Enums(1,2,3) default(1)
// @Param timezone header string false "Time zone" default(Asia/Kolkata)
// @Param offset header integer false "Time zone offset" default(0)
// @Param accept-language header string false "Language: en, hi" Enums(en,hi) default(en)
// @Param appversion header string false "App version" default(v1)
// @Param routeversion header string false "Route version" default(v1)
// @Security basicAuth
// @Success 201 {object} config.APIResponse
// @Failure 400 {object} config.APIResponse
// @Failure 401 {object} config.APIResponse
// @Router /users [post]
func Create(c *fiber.Ctx) error {
	lang := getLang(c)

	var userData CreateUserDTO
	if err := c.BodyParser(&userData); err != nil {
		errortracker.Track(errortracker.LayerController, "Failed to parse request body", err)
		return c.Status(400).JSON(config.Error("Invalid request body", lang))
	}

	if err := CreateUser(&userData); err != nil {
		errortracker.Track(errortracker.LayerController, "Failed to create user", err)
		return c.Status(500).JSON(config.InternalServerError(lang))
	}

	return c.Status(201).JSON(config.Signup(userData, lang))
}
