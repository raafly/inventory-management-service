package listing

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raafly/invetory-management/helper"
	"gorm.io/gorm"
)

func NewAuthRoutes(app *fiber.App, db *gorm.DB) {
	pass := helper.NewPassword()

	repo := NewUserRepository(db)
	serv := NewUserService(repo, pass)
	handler := NewUserController(serv)

	auth := app.Group("/auth")
	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)
}

func NewItemRoutes(app *fiber.App, db *gorm.DB) {
	repo := NewItemRepository(db)
	serv := NewItemService(repo)
	handler := NewItemController(serv)

	item := app.Group("/items")
	item.Post("/new", handler.Create)
	item.Put("/status", handler.UpdateStatus)
	item.Put("/quantity", handler.UpdateQuantity)
	item.Put("/description", handler.UpadteDescription)
	item.Delete("/delete/:id", handler.Delete)
	item.Get("/id/:id", handler.FindById)
	item.Get("/", handler.FindAll)
}

func NewCategoriesRoutes(app *fiber.App, db *gorm.DB) {
	repo := NewCategoryRepository(db)
	serv := NewCategoryService(repo)
	handler := NewCategoryHandler(serv)

	ct := app.Group("/categories")
	ct.Post("/new", handler.Create)
	ct.Put("/update", handler.Update)
	ct.Get("/", handler.GetAllCategory)
}

func NewHistoryRoutes(app *fiber.App, db *gorm.DB) {
	repo := NewHistoryRepository(db)
	serv := NewHistoryService(repo)
	handler := NewHistoryHandler(serv)

	ht := app.Group("/history")
	ht.Get("/find/:id", handler.FindById)
	ht.Get("/", handler.FindAll)
}