package api

import (
	"time"

	"github.com/aryamaahendra/mesiakad/pkgs/api/handlers"
	"github.com/aryamaahendra/mesiakad/pkgs/api/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"gorm.io/gorm"
)

type API struct {
	app         fiber.Router
	db          *gorm.DB
	middlerware *middleware.Middleware
}

func NewAPI(app *fiber.App, DB *gorm.DB) *API {
	middleware := middleware.New(DB)

	api := app.Group("api/v1", limiter.New(limiter.Config{
		Max:        10,
		Expiration: 60 * time.Second,
	}))

	api.Get("ping", func(c *fiber.Ctx) error {
		return c.SendString("success")
	})

	api.Use(middleware.OnlyJSON)

	return &API{
		app:         api,
		db:          DB,
		middlerware: middleware,
	}
}

func (api *API) Init() {
	api.userRoutes()
}

func (api *API) userRoutes() {
	user := handlers.NewUserHandler(api.db)

	api.app.Post("authorized", user.Authorized)

	auth := api.app.Group("/", api.middlerware.OnlyAuthorized)

	auth.Route("user", func(router fiber.Router) {
		router.Get(":username", user.GetByUsername)
		router.Put(":username", user.Update)
		router.Delete(":username", user.Delete)
		router.Get("", user.GetAll)
		router.Post("", user.Create)
	})

	auth.Route("profile", func(router fiber.Router) {
		router.Get("", user.GetProfile)
		router.Put("", user.UpdateProfile)
	})

}
