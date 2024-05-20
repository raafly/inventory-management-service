package listing

import (
	"context"
	"time"

	"github.com/raafly/invetory-management/helper"
)

type UserService interface {
	new(req *Register) error
	login(req *Login) (string, error)
}

type UserServiceImpl struct {
	port UserRepository
	pass *helper.Password
}

func NewUserService(port UserRepository, pass *helper.Password) UserService {
	return &UserServiceImpl{
		port: port,
	}
}

func (s *UserServiceImpl) new(req *Register) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req.Password = s.pass.HashPassword(req.Password)

	err := s.port.new(ctx, req)
	if err != nil {
		return helper.NewInternalServerError()
	}

	return nil
}

func (s *UserServiceImpl) login(req *Login) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := s.port.findByEmail(ctx, req)
	if err != nil {
		return "", helper.NewNotFoundError("account not founded")
	}

	if s.pass.CompareHashAndPassword(resp.Password, req.Password) != nil {
		return "", helper.NewBadRequestError("incorect email and password")
	}

	return resp.Username, nil
}

// // item

type ItemService interface {
	create(req *ReqItem) error
	updateStatus(req *ItemUpdate) error
	updateQuantity(req *ItemUpdate) error
	delete(id int) error
	findById(id int) (*ItemResponse, error)
	FindAll() ([]ItemResponse, error)
	upadteDescription(request *ItemUpdate) error
}

type ItemServiceImpl struct {
	port ItemRepository
}

func NewItemService(port ItemRepository) ItemService {
	return &ItemServiceImpl{
		port: port,
	}
}

func (s *ItemServiceImpl) create(req *ReqItem) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if s.port.new(ctx, req) != nil {
		return helper.NewInternalServerError()
	}

	return nil
}

func (s *ItemServiceImpl) findById(id int) (*ItemResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := s.port.FindById(ctx, id)
	if err != nil {
		return nil, helper.NewNotFoundError("item not found")
	}

	return &ItemResponse{
		Id:          resp.ID,
		Name:        resp.Name,
		Description: resp.Description,
		Category:    resp.Category,
		Quantity:    resp.Quantity,
		Status:      resp.Status,
	}, nil

}

func (s *ItemServiceImpl) updateStatus(req *ItemUpdate) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if s.port.changeStatus(ctx, req.Id, req.Status, req.Quantity) != nil {
		return helper.NewNotFoundError("id item not found")
	}

	return nil
}

func (s *ItemServiceImpl) updateQuantity(req *ItemUpdate) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.port.updateQuantity(ctx, req.Id, req.Quantity); err != nil {
		return helper.NewNotFoundError("id item not found")
	}

	return nil
}

func (s *ItemServiceImpl) delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.port.delete(ctx, id); err != nil {
		return helper.NewNotFoundError("id item not found")
	}

	return nil
}

func (s *ItemServiceImpl) FindAll() ([]ItemResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	items, err := s.port.FindAll(ctx)
	if err != nil {
		return nil, helper.NewInternalServerError()
	}

	var itemResponse []ItemResponse
	for _, item := range items {
		it := ItemResponse{
			Id:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			Quantity:    item.Quantity,
			Category:    item.Category,
			Status:      item.Status,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		}

		itemResponse = append(itemResponse, ItemResponse(it))
	}

	return itemResponse, nil
}

func (s *ItemServiceImpl) upadteDescription(req *ItemUpdate) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.port.addDesc(ctx, req.Id, req.Description); err != nil {
		return helper.NewNotFoundError("id item not found")
	}

	return nil
}

// category

type CategoryService interface {
	save(req *CategoryNew) error
	update(request *CategoryUpdate) error
	getAllCategory() ([]CategoryResponse, error)
}

type CategoryServiceImpl struct {
	Port CategoryRepository
}

func NewCategoryService(port CategoryRepository) CategoryService {
	return &CategoryServiceImpl{
		Port: port,
	}
}

func (s *CategoryServiceImpl) save(req *CategoryNew) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Port.new(ctx, req); err != nil {
		return helper.NewBadRequestError("failed insert data")
	}

	return nil
}

func (s *CategoryServiceImpl) update(req *CategoryUpdate) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Port.update(ctx, req); err != nil {
		return helper.NewBadRequestError("failed insert")
	}

	return nil
}

func (s *CategoryServiceImpl) getAllCategory() ([]CategoryResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	categories, err := s.Port.getAllCategory(ctx)
	if err != nil {
		return nil, helper.NewInternalServerError()
	}

	var category []CategoryResponse
	for _, cate := range categories {
		ct := CategoryResponse{
			ID:          cate.ID,
			Name:        cate.Name,
			Description: cate.Description,
		}

		category = append(category, ct)
	}

	return category, nil
}

type HistoryService interface {
	findById(id int) (*HistoryResponse, error)
	findAll() ([]HistoryResponse, error)
}

type historyService struct {
	Port HistoryRepository
}

func NewHistoryService(port HistoryRepository) HistoryService {
	return &historyService{Port: port}
}

func (s historyService) findById(id int) (*HistoryResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	history, err := s.Port.findById(ctx, id)
	if err != nil {
		return nil, helper.NewNotFoundError("history not founded")
	}

	return &HistoryResponse{
		Id:        history.ID,
		ItemId:    history.ItemID,
		Action:    history.Action,
		Quantity:  history.Quantity,
		UpdatedAt: history.UpdatedAt,
	}, nil
}

func (s historyService) findAll() ([]HistoryResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	histories, err := s.Port.findAll(ctx)
	if err != nil {
		return nil, helper.NewInternalServerError()
	}

	var history []HistoryResponse
	for _, cate := range histories {
		ct := HistoryResponse{
			Id:        cate.ID,
			ItemId:    cate.ItemID,
			Action:    cate.Action,
			Quantity:  cate.Quantity,
			UpdatedAt: cate.UpdatedAt,
		}

		history = append(history, ct)
	}

	return history, nil
}
