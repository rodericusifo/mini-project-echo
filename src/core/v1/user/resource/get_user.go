package resource_user_v1

import (
	"github.com/rodericusifo/mini-project-echo/libs/types"
	"github.com/rodericusifo/mini-project-echo/src/model"
)

func (r *UserResource) GetUser(query *types.Query, payload *model.User) (*model.User, error) {
	return r.Postgres.UserRepository.GetUser(query, payload)
}
