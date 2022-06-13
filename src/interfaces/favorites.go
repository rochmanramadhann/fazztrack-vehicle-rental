package interfaces

import (
	"github.com/rochmanramadhann/fazztrack-vehicle/src/database/gorm/models"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/helpers"
)

type FavoriteRepository interface {
	SearchFavorites(favorite *models.Favorite) (*models.Favorites, error)
	SortFavorites(sort string) (*models.Favorites, error)
	GetFavorites() (*models.Favorites, error)
	GetFavorite(id int) (*models.Favorite, error)
	AddFavorite(favorite *models.Favorite) (*models.Favorite, error)
	UpdateFavorite(favorite *models.Favorite) (*models.Favorite, error)
	DeleteFavorite(favorite *models.Favorite) error
}

type FavoriteService interface {
	SearchFavorites(favorite *models.Favorite) (*helpers.Response, error)
	SortFavorites(sort string) (*helpers.Response, error)
	GetFavorites() (*helpers.Response, error)
	GetFavorite(id int) (*helpers.Response, error)
	AddFavorite(favorite *models.Favorite) (*helpers.Response, error)
	UpdateFavorite(favorite *models.Favorite) (*helpers.Response, error)
	DeleteFavorite(favorite *models.Favorite) (*helpers.Response, error)
}
