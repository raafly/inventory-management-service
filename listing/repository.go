package listing

import (
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	new(ctx context.Context, dto *Register) error
	findByEmail(ctx context.Context, dto *Login) (*User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) new(ctx context.Context, dto *Register) error {
	return r.db.WithContext(ctx).Create(&dto).Error
}

func (r *UserRepositoryImpl) findByEmail(ctx context.Context, dto *Login) (*User, error) {
	var user User
	err := r.db.WithContext(ctx).Where("email = ?", dto.Email).Take(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// item

type ItemRepository interface {
	new(ctx context.Context, item *ReqItem) error
	changeStatus(ctx context.Context, id int, status bool, quantity int) error
	updateQuantity(ctx context.Context, id, quatity int) error
	addDesc(ctx context.Context, id int, desc string) error
	delete(ctx context.Context, itemId int) error
	FindById(ctx context.Context, itemId int) (*Item, error)
	FindAll(ctx context.Context) ([]Item, error)
}

type ItemRepositoryImpl struct {
	db *gorm.DB
}

func NewItemRepository(DB *gorm.DB) ItemRepository {
	return &ItemRepositoryImpl{
		db: DB,
	}
}

func (r ItemRepositoryImpl) new(ctx context.Context, item *ReqItem) error {
	return r.db.WithContext(ctx).Create(&item).Error
}

func (r ItemRepositoryImpl) changeStatus(ctx context.Context, id int, status bool, quantity int) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Update("status", status).Update("quantity", quantity).Error
}

func (r ItemRepositoryImpl) updateQuantity(ctx context.Context, id, quatity int) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Update("quantity", quatity).Error
}

func (r ItemRepositoryImpl) addDesc(ctx context.Context, id int, desc string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Update("description", desc).Error
}

func (r ItemRepositoryImpl) delete(ctx context.Context, id int) error { 
	return r.db.WithContext(ctx).Delete(id, "id = ?").Error
}

func (r ItemRepositoryImpl) FindById(ctx context.Context, id int) (*Item, error) {
	var item Item
	err := r.db.WithContext(ctx).Where("id = ?", id).Take(&item).Error
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r ItemRepositoryImpl) FindAll(ctx context.Context) ([]Item, error) {
	var items []Item
	err := r.db.WithContext(ctx).Table("items").Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items, nil
}

// category

type CategoryRepository interface {
	new(ctx context.Context, dto *CategoryNew) error
	update(ctx context.Context, dto *CategoryUpdate) error
	getAllCategory(ctx context.Context) ([]Category, error)
}

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(Db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{db: Db}
}

func (r CategoryRepositoryImpl) new(ctx context.Context, dto *CategoryNew) error {
	return r.db.WithContext(ctx).Create(&dto).Error
}

func (r CategoryRepositoryImpl) update(ctx context.Context, dto *CategoryUpdate) error {
	return r.db.WithContext(ctx).Where("id = ?", dto.Id).Update("description", dto.Description).Error
}

func (r CategoryRepositoryImpl) getAllCategory(ctx context.Context) ([]Category, error) {
	var category []Category
	err := r.db.WithContext(ctx).Table("category").Find(&category).Error
	if err != nil {
		return nil, err
	}

	return category, nil
}

type HistoryRepository interface {
	findById(ctx context.Context, id int) (*History, error)
	findAll(ctx context.Context) ([]History, error)
}

type historyRepository struct {
	db *gorm.DB
}

func NewHistoryRepository(Db *gorm.DB) HistoryRepository {
	return &historyRepository{db: Db}
}

func (r historyRepository) findById(ctx context.Context, id int) (*History, error) {
	var history History
	err := r.db.WithContext(ctx).Where("id = ?", id).Take(&history).Error
	if err != nil {
		return nil, err
	}

	return &history, nil
}

func (r historyRepository) findAll(ctx context.Context) ([]History, error) {
	var histories []History
	err := r.db.WithContext(ctx).Table("histories").Find(&histories).Error
	if err != nil {
		return nil, err
	}

	return histories, nil
}
