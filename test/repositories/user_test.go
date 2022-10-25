package repositories_test

import (
	"context"
	"testing"

	user "github.com/adrisongomez/project_universidad/repositories"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
    assert := assert.New(t)
	ctx := context.Background()
	userRepository, _ := user.New(ctx)
	payload := user.NewCreateUserPayload(
		"test1",
		"test2@test.com",
		"test1test1",
	)
	createdUser, error := userRepository.CreateRecord(*payload)

	if assert.NotNil(createdUser) {
		assert.Equal(createdUser.Name, payload.GetName())
        assert.NotNil(createdUser.ID)
	} else {
        assert.NotNil(error)
    }

}
