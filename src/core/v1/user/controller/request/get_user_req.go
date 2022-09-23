package request_controller_user_v1

type GetUserRequestParam struct {
	ID string `param:"id" validate:"required,uuid4"`
}

func (r *GetUserRequestParam) CustomValidate() error {
	return nil
}
