//go:build debug

package store

import (
	"github.com/amirkhaki/cnbzdtr/config"
	"github.com/amirkhaki/cnbzdtr/entity"
	"github.com/amirkhaki/cnbzdtr/protocol"

	"context"
	"fmt"
)

type inmemoryStore struct {
	store map[string]*entity.User
}

func (in *inmemoryStore) AddUser(_ context.Context, u *entity.User) error {
	if u.ID == "" {
		return fmt.Errorf("User ID is empty!")
	}
	in.store[u.ID] = u
	return nil
}

func (in *inmemoryStore) GetUserByID(_ context.Context, id string) (*entity.User, error) {
	u, ok := in.store[id]
	if !ok {
		return nil, fmt.Errorf("User not found!")
	}
	return u, nil
}

func (in *inmemoryStore) GetUserOrCreate(_ context.Context, id string) (*entity.User, error) {
	if u, err := in.GetUserByID(id); err == nil {
		return u, nil
	}
	u, err := entity.NewUser(id)
	if err != nil {
		return nil, fmt.Errorf("Could not create user: %w", err)
	}
	err = in.AddUser(u)
	if err != nil {
		return nil, fmt.Errorf("Could not add user: %w", err)
	}
	return u, nil
}

func (in *inmemoryStore) UpdateUser(_ context.Context, u *entity.User) error {
	if u.ID == "" {
		return fmt.Errorf("User ID is empty!")
	}
	if _, ok := in.store[u.ID]; !ok {
		return fmt.Errorf("User not found!")
	}
	in.store[u.ID] = u
	return nil
}

func (in *inmemoryStore) DeleteUser(_ context.Context, u *entity.User) error {
	if u.ID == "" {
		return fmt.Errorf("User ID is empty!")
	}
	if _, ok := in.store[u.ID]; !ok {
		return fmt.Errorf("User not found!")
	}
	delete(in.store, u.ID)
	return nil
}

func New(_ config.Config) (protocol.Store, error) {
	imS := inmemoryStore{}
	imS.store = make(map[string]*entity.User)
	return &imS, nil
}
