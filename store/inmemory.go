//go:build debug

package store

import (
	"github.com/amirkhaki/cnbzdtr/protocol"
	"github.com/amirkhaki/cnbzdtr/entity"

	"fmt"
)
type inmemoryStore struct {
	store map[string]*entity.User
}

func (in *inmemoryStore) AddUser(u *entity.User) error {
	if u.ID == "" {
		return fmt.Errorf("User ID is empty!")
	}
	in.store[u.ID] = u
	return nil
}

func (in *inmemoryStore) GetUserByID(id string) (*entity.User, error) {
	u, ok := in.store[id]
	if !ok {
		return nil, fmt.Errorf("User not found!")
	}
	return u, nil
}


func (in *inmemoryStore) UpdateUser(u *entity.User) error {
	if u.ID == "" {
		return fmt.Errorf("User ID is empty!")
	}
	if _, ok := in.store[u.ID]; !ok {
		return fmt.Errorf("User not found!")
	}
	in.store[u.ID] = u
	return nil
}


func (in *inmemoryStore) DeleteUser(u *entity.User) error {
	if u.ID == "" {
		return fmt.Errorf("User ID is empty!")
	}
	if _, ok := in.store[u.ID]; !ok {
		return fmt.Errorf("User not found!")
	}
	delete(in.store, u.ID)
	return nil
}

func New() protocol.Store {
	imS := inmemoryStore{}
	imS.store = make(map[string]*entity.User)
	return &imS
}
