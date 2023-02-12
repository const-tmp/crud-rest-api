package server

import (
	crud2 "crud-api/crud"
	"crud-api/pkg/common"
	"crud-api/pkg/server"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		AccountID uint
	}

	impl struct {
		crud *server.CRUD[crud2.User, User]
	}
)

func New(db *gorm.DB) crud2.ServerInterface {
	return &impl{crud: server.New[crud2.User, User](
		db,
		func(from *crud2.User) (*User, error) {
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
		func(from *User) (*crud2.User, error) {
			model := common.EncodeModel(from.Model)
			return &crud2.User{
				Id:        model.Id,
				CreatedAt: model.CreatedAt,
				UpdatedAt: model.UpdatedAt,
				DeletedAt: model.DeletedAt,
				AccountId: uint64(from.AccountID),
			}, nil
		},
	)}
}

func (i impl) GetUsers(ctx echo.Context, params crud2.GetUsersParams) error {
	return i.crud.Get(ctx, server.GetParams{
		Offset: params.Offset,
		Limit:  params.Limit,
		Sort:   (*map[string]string)(params.Sort),
	})
}

func (i impl) PostUsers(ctx echo.Context) error {
	return i.crud.Post(ctx)
}

func (i impl) DeleteUsersId(ctx echo.Context, id crud2.Id) error {
	return i.crud.DeleteID(ctx, id)
}

func (i impl) GetUsersId(ctx echo.Context, id crud2.Id) error {
	return i.crud.GetByID(ctx, id)
}

func (i impl) PatchUsersId(ctx echo.Context, id crud2.Id) error {
	return i.crud.PatchByID(ctx, id)
}

func (i impl) PutUsersId(ctx echo.Context, id crud2.Id) error {
	return i.crud.PutByID(ctx, id)
}
