package main

// Mocks

type MockRepo struct{}

func NewMock() *MockRepo {
	return &MockRepo{}
}

func (m *MockRepo) CreateUser() error {
	return nil
}

func (m *MockRepo) CreateTask(t *Task) (*Task, error) {
	return &Task{}, nil
}

func (m *MockRepo) GetTask(id string) (*Task, error) {
	return &Task{}, nil
}
