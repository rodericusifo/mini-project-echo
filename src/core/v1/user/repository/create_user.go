package repository_user_v1

import (
	"github.com/rodericusifo/mini-project-echo/src/model"
)

func (r *UserRepository) CreateUser(payload *model.User) error {
	model := new(model.User)

	sql := r.db.Model(model)

	if payload != nil {
		model = payload
	}

	if err := sql.Create(model).Error; err != nil {
		return err
	}

	return nil
}
