package listing

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/raafly/invetory-management/helper"
)

type UserHandler interface{
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type UserHandlerImpl struct{
	port 	UserService
}

func NewUserController(port UserService) *UserHandlerImpl {
	return &UserHandlerImpl{
		port: port,
	}
}

func (h *UserHandlerImpl) Register(c *fiber.Ctx) error {
	req := new(Register)
	c.BodyParser(req)

	err := h.port.new(req)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	return c.Status(201).JSON(helper.NewCreated("success create account", nil))
}

func (h *UserHandlerImpl) Login(c *fiber.Ctx) error {	
	req := new(Login)
	c.BodyParser(req)

	_, err := h.port.login(req)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	return c.Status(201).JSON(helper.NewSuccess("success login"))
}


// item 
type ItemHandler interface {
	Create(c *fiber.Ctx) error
	UpdateStatus(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	UpdateQuantity(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	UpadteDescription(c *fiber.Ctx) error
}

type ItemHandlerImpl struct {
	port ItemService
}

func NewItemController(port ItemService) *ItemHandlerImpl {
	return &ItemHandlerImpl{
		port: port,
	}
}

func (h *ItemHandlerImpl) Create(c *fiber.Ctx) error {
	req := new(ReqItem)
	c.BodyParser(req)

	err := h.port.create(req)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	return nil
}

func (h *ItemHandlerImpl) UpdateStatus(c *fiber.Ctx) error {
	req := new(ItemUpdate)
	c.BodyParser(req)

	if h.port.updateStatus(req) != nil {
		return c.Status(404).JSON(helper.NewNotFoundError("id not found"))
	}

	return c.Status(200).JSON(helper.NewSuccess("success update item"))
}

func (h *ItemHandlerImpl) UpdateQuantity(c *fiber.Ctx) error {
	req := new(ItemUpdate)
	c.BodyParser(req)

	err := h.port.updateQuantity(req)
	if err != nil {
		return c.Status(404).JSON(err)
	}

	return c.Status(200).JSON(helper.NewSuccess("success update item"))
}

func (h *ItemHandlerImpl) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	new, _ := strconv.Atoi(id)

	err := h.port.delete(new)
	if err != nil {
		return c.Status(404).JSON(err)
	}

	return c.Status(200).JSON(helper.NewSuccess("success delete item"))
}

func (h *ItemHandlerImpl) FindById(c *fiber.Ctx) error {
	new := c.Params("id")
	id, _ := strconv.Atoi(new)
	
	resp, err := h.port.findById(id)
	if err != nil {
		return c.Status(404).JSON(err)
	}

	return c.Status(200).JSON(helper.NewCreated("item found", resp))
}

func (h *ItemHandlerImpl) FindAll(c *fiber.Ctx) error {
	resp, err := h.port.FindAll()
	if err != nil {
		return c.Status(500).JSON(err)
	}

	return c.Status(201).JSON(helper.NewCreated("success", resp))
}

func (h *ItemHandlerImpl) UpadteDescription(c *fiber.Ctx) error {
	req := new(ItemUpdate)
	err := h.port.upadteDescription(req)
	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.Status(200).JSON(helper.NewSuccess("success update desc"))
}

// category

type CategoryHandler interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	GetAllCategory(c *fiber.Ctx) error
}

type CategoryHandlerImpl struct {
	Port CategoryService
}

func NewCategoryHandler(port CategoryService) CategoryHandler {
	return &CategoryHandlerImpl{
		Port: port,
	}
}

func (h *CategoryHandlerImpl) Create(c *fiber.Ctx) error {
	req := new(CategoryNew)

	err := h.Port.save(req)
	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.Status(201).JSON(helper.NewCreated("success create new categories", nil))
}

func (h *CategoryHandlerImpl) Update(c *fiber.Ctx) error {
	req := new(CategoryUpdate)

	err := h.Port.update(req)
	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.Status(200).JSON(helper.NewSuccess("Success update"))
}	

func (h *CategoryHandlerImpl) GetAllCategory(c *fiber.Ctx) error {
	resp, err := h.Port.getAllCategory()
	if err != nil {
		return c.Status(500).JSON(err)
	}

	return c.Status(201).JSON(helper.NewCreated("success", resp))
}

type HistoryHandler interface {
	FindById(c *fiber.Ctx) error	
	FindAll(c *fiber.Ctx) error
}

type historyHandler struct {
	Port HistoryService
}

func NewHistoryHandler(port HistoryService) HistoryHandler {
	return &historyHandler{Port: port}
}

func (h *historyHandler) FindById(c *fiber.Ctx) error {
	new := c.Params("id")
	id, _ := strconv.Atoi(new)
	
	resp, err := h.Port.findById(id)
	if err != nil {
		return c.Status(404).JSON(err)
	}

	return c.Status(200).JSON(helper.NewCreated("success", resp))

}
func (h *historyHandler) FindAll(c *fiber.Ctx) error {
	resp, err := h.Port.findAll()
	if err != nil {
		return c.Status(500).JSON(err)
	}

	return c.Status(200).JSON(helper.NewCreated("success", resp))
}