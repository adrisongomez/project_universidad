package bases

import (
	"context"
	"fmt"

	"github.com/adrisongomez/project_universidad/models"
)

type BaseRepository struct {
	Client  *models.PrismaClient
	Context context.Context
}

type IRepository[T BaseRepository, C any, U any] interface {
	FindRecordById(id string) (T, error)
	CreateRecord(id string, data C) (T, error)
	UpdateRecord(id string, data U) (T, error)
}

func NewBaseRepository(ctx context.Context) (*BaseRepository, *RepositoryError) {
	client := models.NewClient()
	err := client.Prisma.Connect()
	if err != nil {
		return nil, NewRepositoryError("Fail trying to connect with database", "CONNECT_DB_ERROR")
	}

	return &BaseRepository{
		Context: ctx,
		Client:  client,
	}, nil
}

func NewRepositoryError(message string, Type string) *RepositoryError {
    return &RepositoryError{
        Message: message,
        Type: Type,
    }
}

type RepositoryError struct {
	Message string
	Type    string
}

func (e *RepositoryError) Error() string {
    return fmt.Sprintf("%v, %v", e.Message, e.Type)
}

