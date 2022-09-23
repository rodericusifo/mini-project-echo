package request_controller_user_v1

import "time"

type CreateUserRequestBody struct {
	Name        string     `json:"name" validate:"required"`
	Email       string     `json:"email" validate:"required,email"`
	Description *string    `json:"description,omitempty" validate:"omitempty"`
	Age         *int64     `json:"age,omitempty" validate:"omitempty,min=0"`
	Birthday    *time.Time `json:"birthday,omitempty" validate:"omitempty"`
}

func (r *CreateUserRequestBody) CustomValidate() error {
	return nil
}
