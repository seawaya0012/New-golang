package routes

import (
	"backend/controllers"

	// "github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func UserRoute(app *fiber.App) {

	api := app.Group("/api") // /api
	v1 := api.Group("/v1")   // /api/v1
	v2 := api.Group("/v2")   // /api/v2

	v1.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"john":   "doe",
			"admin":  "123456",
			"testgo": "772565",
		},
	}))

	v1.Get("/hello", controllers.HelloTest)
	v1.Post("/", controllers.HelloTest1)
	v1.Get("/user/:name", controllers.HelloTest2)
	v1.Post("/register/", controllers.HelloTest3)

	v1.Post("/dog", controllers.AddDog)
	v1.Get("/dog", controllers.GetDogs)
	v1.Get("/dog/filter", controllers.GetDog)
	v1.Put("/dog/:id", controllers.UpdateDog)
	v1.Delete("/dog/:id", controllers.RemoveDog)

	v1.Post("/work/calculate_factorial/:id", controllers.AddCalculate)
	v1.Post("/work/register/", controllers.Register)
	v1.Get("/work/registers/", controllers.GetID)
	v1.Delete("/work/register/:id", controllers.DeleteID)

	v1.Get("/dog/jason", controllers.GetDogsJson)

	//assign on 1/9/65

	v1.Post("/profile", controllers.AddProfile)
	v2.Get("/profile", controllers.GetProfile)
	v1.Put("/profile/:id", controllers.UpdateProfile)
	v1.Delete("/profile/:id", controllers.RemoveProfile)

	app.Listen(":3000")
}
