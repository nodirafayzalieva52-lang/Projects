package storage

import (
	"encoding/json"
	"errors"
	"os"
	"project/models"
	"sync"
)

type UserStorage struct {
	mu       sync.Mutex
	filename string
}

func New(filename string) *UserStorage {
		return &UserStorage{
		filename: filename,
	}
}

func (s *UserStorage) GetAll() ([]models.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	file, err := os.ReadFile(s.filename)
	if err != nil {
		return nil, err
	}
	var users []models.User
	err = json.Unmarshal(file, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserStorage) GetByID(id int) (*models.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	file, err := os.ReadFile(s.filename)
	if err != nil {
		return nil, err
	}
	
	var users []models.User
	err = json.Unmarshal(file, &users)
	if err != nil {
		return nil, err
	}
	for _, val := range users{
		if val.ID == id{
			return &val,nil
		}
	}

	return nil,errors.New("user not found")
}

func (s *UserStorage) Create(user models.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	file, err := os.ReadFile(s.filename)
	if err != nil {
		return err
	}
	
	var users []models.User
	err = json.Unmarshal(file, &users)
	if err != nil {
		return err
	}
	users = append(users, user)

	file1, err := json.Marshal(users)
	os.WriteFile("data/users.json", file1, 0644)
	return nil
}

func (s *UserStorage) Update(id int, user models.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	file, err := os.ReadFile(s.filename)
	if err != nil {
		return err
	}

	var users []models.User
	err = json.Unmarshal(file, &users)
	if err != nil {
		return err
	}

	isUpdated := false

	for i, val := range users {
		if val.ID == id {
			isUpdated = true
			users[i] = user
			break
		}
	}

	if isUpdated {
		return nil
	}
	return errors.New("user not found")
}
