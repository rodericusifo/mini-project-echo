package user_postgres_repository

import (
	"github.com/rodericusifo/mini-project-echo/libs/constant"
	"github.com/rodericusifo/mini-project-echo/libs/types"
	"github.com/rodericusifo/mini-project-echo/libs/util"
	"github.com/rodericusifo/mini-project-echo/src/model"
)

func (r *UserRepository) GetUser(query *types.Query, payload *model.User) (*model.User, error) {
	model := new(model.User)

	sql := r.db.Model(model)

	if query != nil {
		if len(query.SelectColumns) > 0 {
			sql = sql.Select(util.UniqueSlice(query.SelectColumns, constant.DEFAULT_SELECT_COLUMNS))
		}
	}

	if payload != nil {
		if payload.ID != "" {
			sql = sql.Where("id = ?", payload.ID)
		}
	}

	if err := sql.First(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}
