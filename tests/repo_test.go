package tests

import "github.com/rryowa/Gojira-project-manager/entity"

// Mocks

type MockRepo struct{}

func (m *MockRepo) CreateUser(u *entity.User) (*entity.User, error) {
	return &entity.User{}, nil
}
func (m *MockRepo) GetUserByID(id string) (*entity.User, error) {
	return &entity.User{}, nil
}
func (m *MockRepo) CreateTask(t *entity.Task) (*entity.Task, error) {
	return &entity.Task{}, nil
}
func (m *MockRepo) GetTask(id string) (*entity.Task, error) {
	return &entity.Task{}, nil
}
