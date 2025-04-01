package createuser

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"servers/internal/models"
	"testing"
)

type mockRepository struct {
	mock.Mock
}

func (m *mockRepository) SaveUser(user models.User) (int, error) {
	args := m.Called(user)
	return args.Int(0), args.Error(1)
}

func TestCreateUser(t *testing.T) {
	// Arrange
	repoMock := &mockRepository{}
	repoMock.On("SaveUser", mock.AnythingOfType("models.User")).Return(2, nil)
	uc := NewUserUseCase(repoMock)

	// Act
	id, err := uc.CreateUser()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 2, id)
	repoMock.AssertExpectations(t)
}
func TestCreateUser_HTTPError(t *testing.T) {
	// Arrange
	repoMock := &mockRepository{}
	repoMock.On("SaveUser", mock.AnythingOfType("models.User")).Return(-1, errors.New("some error"))
	uc := NewUserUseCase(repoMock)
	uc.apiURL = "http://fakevdsgbhadfndffffffffffffffffffffffurl.com"

	_, err := uc.CreateUser()

	// Assert
	assert.Error(t, err)
	repoMock.AssertNotCalled(t, "SaveUser")
}
