package user

import (
	"context"
	"errors"

	"github.com/adrisongomez/project_universidad/bases"
	"github.com/adrisongomez/project_universidad/models"
    "github.com/adrisongomez/project_universidad/utils"
)


type UserRepository struct {
	*bases.BaseRepository
}

func New(ctx context.Context) (*UserRepository, *bases.RepositoryError) {
	repository, err := bases.NewBaseRepository(ctx)
	if err != nil {
		return nil, err
	}
	return &UserRepository{
		BaseRepository: repository,
	}, nil

}



func (this UserRepository) FindRecordById(id string) *models.UserModel {
	response, _ := this.Client.User.FindUnique(
		models.User.ID.Equals(id),
	).Exec(this.Context)

	return response
}

func (this UserRepository) CreateRecord(data CreateUser) (*models.UserModel, error) {
	password_hash, err := password.GeneratePasswordHash(data.password)
	if err != nil {
		return nil, errors.New("Fail hashing the user password")
	}
	response, err := this.Client.User.CreateOne(
		models.User.Name.Set(data.name),
		models.User.Email.Set(data.email),
		models.User.Password.Set(*password_hash),
	).Exec(this.Context)

	if err != nil {
		return nil, errors.New("Fail creating user record on DB")
	}
	return response, nil
}

type UpdateUser struct {
	name     string
	email    string
	password string
}

func (this UserRepository) UpdateRecord(id string, data UpdateUser) *models.UserModel {
	response, _ := this.Client.User.FindUnique(
		models.User.ID.Equals(id),
	).Update(
		models.User.Name.SetIfPresent(&data.name),
		models.User.Name.SetIfPresent(&data.email),
		models.User.Name.SetIfPresent(&data.password),
	).Exec(this.Context)

	return response
}

