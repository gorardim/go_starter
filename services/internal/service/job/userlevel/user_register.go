package userlevel

import (
	"context"

	"app/api/job"
	"app/pkg/errx"
	"app/services/internal/repo"
)

var _ job.UserRegisterConsumer = (*UserRegister)(nil)

type UserRegister struct {
	UserRepo *repo.UserRepo
}

func (u *UserRegister) Default(ctx context.Context, request *job.UserRegisterRequest) error {
	// query user
	_, err := u.UserRepo.FindById(ctx, request.UserId)
	if err != nil {
		return errx.FilterRecordNotFoundErr(err)
	}
	return nil
}
