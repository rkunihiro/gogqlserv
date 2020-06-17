package repos

import (
	"github.com/jinzhu/gorm"
	"github.com/rkunihiro/gogqlserv/entity"
)

// UserRepo User entity repository
type UserRepo interface {
	Register(name string) (*entity.User, error)
	FindByID(id entity.UserID) (*entity.User, error)
	FindAll() ([]*entity.User, error)
}

// NewUserRepo generate new UserRepo implements
func NewUserRepo(db *gorm.DB) (UserRepo, error) {
	return &userRepo{
		db: db,
	}, nil
}

type userRepo struct {
	db *gorm.DB
}

func (p *userRepo) Register(name string) (*entity.User, error) {
	user := &entity.User{
		Name: name,
	}
	err := p.db.Save(user).Error
	return user, err
}

func (p *userRepo) FindByID(id entity.UserID) (*entity.User, error) {
	user := &entity.User{}
	result := p.db.Where(&entity.User{ID: id}).First(user)
	if result.RecordNotFound() {
		return nil, nil
	}
	return user, result.Error
}

func (p *userRepo) FindAll() ([]*entity.User, error) {
	users := make([]*entity.User, 0)
	err := p.db.Find(&users).Error
	return users, err
}
