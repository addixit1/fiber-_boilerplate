package main

import (
	"log"

	"github.com/addixit1/fiber-boilerplate/internal/app"
)

// @title  Fiber Boilerplate
// @version 1.0
// @description Production ready Fiber backend
// @termsOfService https://example.com/terms

// @contact.name Aman Dixit
// @contact.email aman.dixit@appinevtiv.com

// @host localhost:3010
// @BasePath /api/v1

// 🔐 Auth
// @securityDefinitions.basic BasicAuth
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// 📋 Global Headers
// @Param platform header string false "Device OS: 1-Android, 2-iOS, 3-WEB" Enums(1, 2, 3) default(1)
// @Param timezone header string false "Time zone" default(Asia/Kolkata)
// @Param offset header number false "Time zone offset" default(0)
// @Param accept-language header string false "Language preference: en, hi" Enums(en, hi) default(en)
// @Param appversion header string false "App version" default(v1)
// @Param routeversion header string false "Route version" default(v1)

func main() {
	fiberApp := app.New()
	log.Fatal(fiberApp.Listen(":3010"))

}
