package users

import (
	"fmt"
	"strings"

	"github.com/rochmanramadhann/fazztrack-vehicle/src/database/gorm/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) SearchUsers(user *models.User) (*models.Users, error) {
	var users models.Users

	result := r.db.Where("name = ?", user.Name).Or("email = ?", user.Email).Or("password = ?", user.Password).Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return &users, nil
}

func (r *UserRepository) SortUsers(sort string) (*models.Users, error) {
	var users models.Users

	var result *gorm.DB

	if strings.ToLower(sort) == "asc" {
		result = r.db.Order("id asc").Find(&users)
	} else if strings.ToLower(sort) == "desc" {
		result = r.db.Order("id desc").Find(&users)
	} else {
		result = r.db.Find(&users)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &users, nil
}

func (r *UserRepository) GetUsers() (*models.Users, error) {
	var users models.Users

	result := r.db.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return &users, nil
}

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User

	result := r.db.First(&user, "name = ?", username)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *UserRepository) GetUser(id int) (*models.User, error) {
	var user models.User

	result := r.db.First(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *UserRepository) AddUser(user *models.User) (*models.User, error) {
	result := r.db.Create(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *UserRepository) UpdateUser(user *models.User) (*models.User, error) {
	// userExist := r.db.First(user, user.Model.ID)
	// if userExist.Error != nil {
	// 	fmt.Println("1")
	// 	return nil, userExist.Error
	// }
	result := r.db.Save(user)

	if result.Error != nil {
		fmt.Println("2")
		return nil, result.Error
	}

	fmt.Println("3")
	return user, nil
}

func (r *UserRepository) DeleteUser(user *models.User) error {
	userExist := r.db.First(&user, user.Model.ID)
	if userExist.Error != nil {
		fmt.Println("1")
		return userExist.Error
	}
	result := r.db.Delete(&user)

	if result.Error != nil {
		fmt.Println("2")
		return result.Error
	}

	fmt.Println("3")
	return nil
}
