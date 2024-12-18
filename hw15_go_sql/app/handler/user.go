package handler

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kotoproger/home_work_basic/hw15_go_sql/app"
	"github.com/kotoproger/home_work_basic/hw15_go_sql/app/security"
	"github.com/kotoproger/home_work_basic/hw15_go_sql/internal/repository"
	"github.com/kotoproger/home_work_basic/hw15_go_sql/internal/repositorywrapper"
)

type User struct {
	app app.App
}

func NewUser(a app.App) *User {
	return &User{app: a}
}

type UserDto struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	password     string `json:"-"`
	passwordHash string `json:"-"`
	passwordSalt string `json:"-"`
}

func NewUserDtoWithPassword(name string, email string, password string) (*UserDto, error) {
	salt := security.GenSalt()
	hash, err := security.HashPassword(password, salt)
	if err != nil {
		return nil, fmt.Errorf("create user dto: %w", err)
	}
	return &UserDto{
		Name:         name,
		Email:        email,
		password:     password,
		passwordSalt: salt,
		passwordHash: *hash,
	}, nil
}

func (u *User) Auth(ctx *gin.Context) {
	userID := ctx.Request.Header.Get("Auth")
	if userID != "" {
		ctx.Set("userID", userID)
	}
	ctx.Next()
}

func (u *User) Register(ctx context.Context, user UserDto) (*UserDto, error) {
	userIDAny, err := u.app.Repository.RunTransactional(ctx, func(repo repository.Querier) (any, error) {
		uuid, err := repo.CreateUser(
			ctx,
			repository.CreateUserParams{
				Name:         &user.Name,
				Email:        user.Email,
				PasswordHash: user.passwordHash,
				PasswordSalt: user.passwordSalt,
			},
		)
		if err != nil {
			return nil, fmt.Errorf("repository create user: %w", err)
		}
		return repositorywrapper.UUIDToString(uuid)
	})
	if err != nil {
		return nil, err
	}

	stringID, ok := userIDAny.(string)
	if !ok {
		return nil, fmt.Errorf("uuid to string conversion")
	}

	user.ID = stringID

	return &user, nil
}
