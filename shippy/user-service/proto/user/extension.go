package go_micro_srv_user

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

func (_ *User) BeforeCreate(scope *gorm.Scope) error {
	u := uuid.NewV4()
	return scope.SetColumn("Id", u.String())
}
