package server

import (
	"github.com/labstack/echo/v4"
	"github.com/nullc4t/crud-rest-api/crud"
	"github.com/nullc4t/crud-rest-api/pkg/common"
	"github.com/nullc4t/crud-rest-api/pkg/repo"
	echoserver "github.com/nullc4t/crud-rest-api/pkg/server/echo"
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		AccountID uint
	}

	impl struct {
		crud *echoserver.Server[crud.User, User]
	}
)

func New(db *gorm.DB) crud.ServerInterface {
	return &impl{crud: echoserver.New[crud.User, User](
		db,
		func(from *crud.User) (*User, error) {
			return &User{
				Model: common.DecodeModel(common.Model{
					Id:        from.Id,
					CreatedAt: from.CreatedAt,
					UpdatedAt: from.UpdatedAt,
					DeletedAt: from.DeletedAt,
				}),
				AccountID: uint(from.AccountId),
			}, nil
		},
		func(from *User) (*crud.User, error) {
			model := common.EncodeModel(from.Model)
			return &crud.User{
				Id:        model.Id,
				CreatedAt: model.CreatedAt,
				UpdatedAt: model.UpdatedAt,
				DeletedAt: model.DeletedAt,
				AccountId: uint64(from.AccountID),
			}, nil
		},
	)}
}

func (i impl) GetUsers(ctx echo.Context, params crud.GetUsersParams) error {
	return i.crud.Get(ctx, echoserver.GetParams{
		Offset: params.Offset,
		Limit:  params.Limit,
		Sort:   (*repo.Sort)(params.Sort),
	})
}

func (i impl) PostUsers(ctx echo.Context) error {
	return i.crud.Post(ctx)
}

func (i impl) DeleteUsersId(ctx echo.Context, id crud.Id) error {
	return i.crud.DeleteID(ctx, id)
}

func (i impl) GetUsersId(ctx echo.Context, id crud.Id) error {
	return i.crud.GetByID(ctx, id)
}

func (i impl) PatchUsersId(ctx echo.Context, id crud.Id) error {
	return i.crud.PatchByID(ctx, id)
}

func (i impl) PutUsersId(ctx echo.Context, id crud.Id) error {
	return i.crud.PutByID(ctx, id)
}
