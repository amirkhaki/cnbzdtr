package protocol

import (
	"context"
	"github.com/amirkhaki/cnbzdtr/entity"
)

type Store interface {
	UserCRUD
}
type UserCRUD interface {
	AddUser(ctx context.Context, u *entity.User) error
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
	GetUserOrCreate(ctx context.Context, id string) (*entity.User, error)
	UpdateUser(ctx context.Context, u *entity.User) error
	DeleteUser(ctx context.Context, u *entity.User) error
}
