package favorites

import (
	"github.com/rochmanramadhann/fazztrack-vehicle/src/database/gorm/models"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/helpers"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/interfaces"
)

type FavoriteService struct {
	repository interfaces.FavoriteRepository
}

func NewFavoriteService(repository interfaces.FavoriteRepository) *FavoriteService {
	return &FavoriteService{repository}
}

func (r *FavoriteService) SearchFavorites(favorite *models.Favorite) (*helpers.Response, error) {
	data, err := r.repository.SearchFavorites(favorite)

	if err != nil {
		response := helpers.NewResponse(data, 400, true, err.Error())
		return response, err
	}

	if len(*data) == 0 {
		response := helpers.NewResponse(data, 404, false, "Favorites Not Found")
		return response, nil
	}

	response := helpers.NewResponse(data, 200, false, "Favorites Found")
	return response, nil
}

func (r *FavoriteService) SortFavorites(sort string) (*helpers.Response, error) {
	data, err := r.repository.SortFavorites(sort)

	if err != nil {
		response := helpers.NewResponse(data, 400, true, err.Error())
		return response, err
	}

	if len(*data) == 0 {
		response := helpers.NewResponse(data, 404, false, "Favorites Not Found")
		return response, nil
	}

	response := helpers.NewResponse(data, 200, false, "Favorites Found")
	return response, nil
}

func (r *FavoriteService) GetFavorites() (*helpers.Response, error) {
	data, err := r.repository.GetFavorites()

	if err != nil {
		response := helpers.NewResponse(data, 400, true, err.Error())
		return response, err
	}

	if len(*data) == 0 {
		response := helpers.NewResponse(data, 404, false, "Favorites Not Found")
		return response, nil
	}

	response := helpers.NewResponse(data, 200, false, "Favorites Found")
	return response, nil
}

func (r *FavoriteService) GetFavorite(id int) (*helpers.Response, error) {
	data, err := r.repository.GetFavorite(id)

	if err != nil {
		response := helpers.NewResponse(data, 400, true, err.Error())
		return response, err
	}

	// if *&data.Model.ID == 0 {
	// 	response := helpers.NewResponse(data, 404, false, "Favorite Not Found")
	// 	return response, nil
	// }

	response := helpers.NewResponse(data, 200, false, "Favorite Found")
	return response, nil
}

func (r *FavoriteService) AddFavorite(favorite *models.Favorite) (*helpers.Response, error) {
	data, err := r.repository.AddFavorite(favorite)

	if err != nil {
		response := helpers.NewResponse(data, 400, true, err.Error())
		return response, err
	}

	// if *&data.Model.ID == 0 {
	// 	response := helpers.NewResponse(data, 404, false, "Favorite Not Found")
	// 	return response, nil
	// }

	response := helpers.NewResponse(data, 200, false, "Success Add Favorite")
	return response, nil
}

func (r *FavoriteService) UpdateFavorite(favorite *models.Favorite) (*helpers.Response, error) {
	data, err := r.repository.UpdateFavorite(favorite)

	if err != nil {
		response := helpers.NewResponse(data, 400, true, err.Error())
		return response, err
	}

	// if *&data.Model.ID == 0 {
	// 	response := helpers.NewResponse(data, 404, false, "Favorite Not Found")
	// 	return response, nil
	// }

	response := helpers.NewResponse(data, 200, false, "Favorite Found")
	return response, nil
}

func (r *FavoriteService) DeleteFavorite(favorite *models.Favorite) (*helpers.Response, error) {
	err := r.repository.DeleteFavorite(favorite)

	if err != nil {
		response := helpers.NewResponse(err, 400, true, err.Error())
		return response, err
	}

	// if *&data.Model.ID == 0 {
	// 	response := helpers.NewResponse(data, 404, false, "Favorite Not Found")
	// 	return response, nil
	// }

	response := helpers.NewResponse(err, 200, false, "Success Delete Favorite")
	return response, nil
}

// func (r *FavoriteService) FindByFavoritename(Favoritename string) (*helpers.Response, error) {
// 	data, err := r.repository.FindByFavoritename(Favoritename)
// 	if err != nil {
// 		return nil, err
// 	}

// 	response := helpers.New(data, 200, false)
// 	return response, nil
// }
// func (r *FavoriteService) Save(usr *models.Favorite) (*helpers.Response, error) {
// 	hassPssword, err := helpers.HashPasword(usr.Password)
// 	if err != nil {
// 		return nil, err
// 	}

// 	usr.Password = hassPssword
// 	data, err := r.repository.Add(usr)
// 	if err != nil {
// 		return nil, err
// 	}

// 	response := helpers.New(data, 200, false)
// 	return response, nil
// }
