package favorites

import (
	"fmt"
	"strings"
	"time"

	"github.com/rochmanramadhann/fazztrack-vehicle/src/database/gorm/models"
	"gorm.io/gorm"
)

type FavoriteRepository struct {
	db *gorm.DB
}

func NewFavoriteRepository(db *gorm.DB) *FavoriteRepository {
	return &FavoriteRepository{db}
}

func (r *FavoriteRepository) SearchFavorites(favorite *models.Favorite) (*models.Favorites, error) {
	var favorites models.Favorites

	result := r.db.Where(" = ?", time.Now()).Find(&favorites)

	if result.Error != nil {
		return nil, result.Error
	}

	return &favorites, nil
}

func (r *FavoriteRepository) SortFavorites(sort string) (*models.Favorites, error) {
	var favorites models.Favorites

	var result *gorm.DB

	if strings.ToLower(sort) == "asc" {
		result = r.db.Order("id asc").Find(&favorites)
	} else if strings.ToLower(sort) == "desc" {
		result = r.db.Order("id desc").Find(&favorites)
	} else {
		result = r.db.Find(&favorites)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &favorites, nil
}

func (r *FavoriteRepository) GetFavorites() (*models.Favorites, error) {
	var favorites models.Favorites

	result := r.db.Find(&favorites)

	if result.Error != nil {
		return nil, result.Error
	}

	return &favorites, nil
}

func (r *FavoriteRepository) GetFavorite(id int) (*models.Favorite, error) {
	var favorite models.Favorite

	result := r.db.First(&favorite, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &favorite, nil
}

func (r *FavoriteRepository) AddFavorite(favorite *models.Favorite) (*models.Favorite, error) {
	result := r.db.Create(favorite)

	if result.Error != nil {
		return nil, result.Error
	}

	return favorite, nil
}

func (r *FavoriteRepository) UpdateFavorite(favorite *models.Favorite) (*models.Favorite, error) {
	// FavoriteExist := r.db.First(Favorite, Favorite.Model.ID)
	// if FavoriteExist.Error != nil {
	// 	fmt.Println("1")
	// 	return nil, FavoriteExist.Error
	// }
	result := r.db.Save(favorite)

	if result.Error != nil {
		fmt.Println("2")
		return nil, result.Error
	}

	fmt.Println("3")
	return favorite, nil
}

func (r *FavoriteRepository) DeleteFavorite(favorite *models.Favorite) error {
	FavoriteExist := r.db.First(&favorite, favorite.Id)
	if FavoriteExist.Error != nil {
		fmt.Println("1")
		return FavoriteExist.Error
	}
	result := r.db.Delete(&favorite)

	if result.Error != nil {
		fmt.Println("2")
		return result.Error
	}

	fmt.Println("3")
	return nil
}
