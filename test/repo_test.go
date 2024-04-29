package test

import "github.com/rryowa/go-jwt-auth/entity"

// Mocks

type MockRepo struct{}

func NewMock() *MockRepo {
	return &MockRepo{}
}

func (m *MockRepo) CreateUser() error {
	return nil
}

func (m *MockRepo) CreateTask(t *entity.Task) (*entity.Task, error) {
	return &entity.Task{}, nil
}

func (m *MockRepo) GetTask(id string) (*entity.Task, error) {
	return &entity.Task{}, nil
}
