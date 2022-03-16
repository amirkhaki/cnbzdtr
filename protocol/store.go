package protocol

import (
	"github.com/amirkhaki/cnbzdtr/entity"
)

type Store interface {
	UserCRUD
}
type UserCRUD interface {
	AddUser(u *entity.User) error
	GetUserByID(id string) (*entity.User, error)
	UpdateUser(u *entity.User) error
	DeleteUser(u *entity.User) error
}
