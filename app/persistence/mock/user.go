package mock

import (
	"sync"

	"github.com/google/uuid"
	"github.com/kkatou7209/godo/app/domain/entity"
	"github.com/kkatou7209/godo/app/domain/value"
	"github.com/kkatou7209/godo/app/port/out/dto"
)

type MockUserRepository struct {
	users map[value.UserId]*entity.User
	mu sync.Mutex
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[value.UserId]*entity.User),
		mu: sync.Mutex{},
	}
}

func (r *MockUserRepository) Create(user *dto.CreateUserCommand) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	u := entity.NewUser(
		value.NewUserId(uuid.NewString()),
		user.UserName,
		user.Email,
		user.Password,
	)

	r.users[u.Id()] = u

	return nil
}

func (r *MockUserRepository) GetById(userId value.UserId) (*entity.User, error) {

	r.mu.Lock()
	defer r.mu.Unlock()

	return r.users[userId], nil
}

func (r *MockUserRepository) GetByEmail(email value.Email) (*entity.User, error) {

	r.mu.Lock()
	defer r.mu.Unlock()

	for _, u := range r.users {

		if u.Email() == email {
			return u, nil
		}
	}

	return nil, nil
}

func (r *MockUserRepository) Update(user *entity.User) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[user.Id()] = user

	return nil
}