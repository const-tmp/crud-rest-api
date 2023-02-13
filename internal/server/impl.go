package server

import (
	"github.com/labstack/echo/v4"
	"github.com/nullc4t/crud-rest-api/pkg/auth"
	"github.com/nullc4t/crud-rest-api/pkg/common"
	"github.com/nullc4t/crud-rest-api/pkg/repo"
	echoserver "github.com/nullc4t/crud-rest-api/pkg/server/echo"
	"gorm.io/gorm"
)

type (
	// Account defines DB model for Account.
	Account struct {
		gorm.Model
		Blocked bool
		Name    string `gorm:"unique"`
	}

	// Permission defines DB model for Permission.
	Permission struct {
		gorm.Model
		Name      string
		AccountID int
		Account   Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		ServiceID int
		Service   Service `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		UserID    int
	}

	// Service defines DB model for Service.
	Service struct {
		gorm.Model
		Name string `gorm:"unique"`
	}

	impl struct {
		account     *echoserver.Server[auth.Account, Account]
		service     *echoserver.Server[auth.Service, Service]
		permissions *echoserver.Server[auth.Permission, Permission]
	}
)

func New(db *gorm.DB) auth.ServerInterface {
	return &impl{
		account:     NewAccountServer(db),
		service:     NewServiceServer(db),
		permissions: NewPermissionServer(db),
	}
}

func NewAccountServer(db *gorm.DB) *echoserver.Server[auth.Account, Account] {
	return echoserver.New[auth.Account, Account](
		db,
		func(from *auth.Account) (*Account, error) {
			var blocked bool
			if from.Blocked != nil && *from.Blocked {
				blocked = true
			}
			return &Account{
				Model: common.DecodeModel(common.Model{
					Id:        from.Id,
					CreatedAt: from.CreatedAt,
					UpdatedAt: from.UpdatedAt,
					DeletedAt: from.DeletedAt,
				}),
				Blocked: blocked,
				Name:    from.Name,
			}, nil
		},
		func(from *Account) (*auth.Account, error) {
			model := common.EncodeModel(from.Model)
			var blocked *bool
			if from.Blocked {
				blocked = new(bool)
				*blocked = true
			}
			return &auth.Account{
				Blocked:   blocked,
				CreatedAt: model.CreatedAt,
				DeletedAt: model.DeletedAt,
				Id:        model.Id,
				Name:      from.Name,
				UpdatedAt: model.UpdatedAt,
			}, nil
		},
	)
}

func NewServiceServer(db *gorm.DB) *echoserver.Server[auth.Service, Service] {
	return echoserver.New[auth.Service, Service](
		db,
		func(from *auth.Service) (*Service, error) {
			return &Service{
				Model: common.DecodeModel(common.Model{
					Id:        from.Id,
					CreatedAt: from.CreatedAt,
					UpdatedAt: from.UpdatedAt,
					DeletedAt: from.DeletedAt,
				}),
				Name: from.Name,
			}, nil
		},
		func(from *Service) (*auth.Service, error) {
			model := common.EncodeModel(from.Model)
			return &auth.Service{
				CreatedAt: model.CreatedAt,
				DeletedAt: model.DeletedAt,
				Id:        model.Id,
				Name:      from.Name,
				UpdatedAt: model.UpdatedAt,
			}, nil
		},
	)
}

func NewPermissionServer(db *gorm.DB) *echoserver.Server[auth.Permission, Permission] {
	return echoserver.New[auth.Permission, Permission](
		db,
		func(from *auth.Permission) (*Permission, error) {
			return &Permission{
				Model: common.DecodeModel(common.Model{
					Id:        from.Id,
					CreatedAt: from.CreatedAt,
					UpdatedAt: from.UpdatedAt,
					DeletedAt: from.DeletedAt,
				}),
				Name:      from.Name,
				AccountID: from.AccountId,
				ServiceID: from.ServiceId,
				UserID:    from.UserId,
			}, nil
		},
		func(from *Permission) (*auth.Permission, error) {
			model := common.EncodeModel(from.Model)
			return &auth.Permission{
				AccountId: from.AccountID,
				CreatedAt: model.CreatedAt,
				DeletedAt: model.DeletedAt,
				Id:        model.Id,
				Name:      from.Name,
				ServiceId: from.ServiceID,
				UpdatedAt: model.UpdatedAt,
				UserId:    from.UserID,
			}, nil
		},
	)
}

func (i impl) GetAccounts(ctx echo.Context, params auth.GetAccountsParams) error {
	return i.account.Get(ctx, echoserver.GetParams{
		Offset: params.Offset,
		Limit:  params.Limit,
		Sort:   (*repo.Sort)(params.Sort),
	})
}

func (i impl) PostAccounts(ctx echo.Context) error {
	return i.account.Post(ctx)
}

func (i impl) DeleteAccountsId(ctx echo.Context, id auth.Id) error {
	return i.account.DeleteID(ctx, id)
}

func (i impl) GetAccountsId(ctx echo.Context, id auth.Id) error {
	return i.account.GetByID(ctx, id)
}

func (i impl) PatchAccountsId(ctx echo.Context, id auth.Id) error {
	return i.account.PatchByID(ctx, id)
}

func (i impl) PutAccountsId(ctx echo.Context, id auth.Id) error {
	return i.account.PutByID(ctx, id)
}

func (i impl) GetPermissions(ctx echo.Context, params auth.GetPermissionsParams) error {
	return i.permissions.Get(ctx, echoserver.GetParams{
		Offset: params.Offset,
		Limit:  params.Limit,
		Sort:   (*repo.Sort)(params.Sort),
	})
}

func (i impl) PostPermissions(ctx echo.Context) error {
	return i.permissions.Post(ctx)
}

func (i impl) DeletePermissionsId(ctx echo.Context, id auth.Id) error {
	return i.permissions.DeleteID(ctx, id)
}

func (i impl) GetPermissionsId(ctx echo.Context, id auth.Id) error {
	return i.permissions.GetByID(ctx, id)
}

func (i impl) PatchPermissionsId(ctx echo.Context, id auth.Id) error {
	return i.permissions.PatchByID(ctx, id)
}

func (i impl) PutPermissionsId(ctx echo.Context, id auth.Id) error {
	return i.permissions.PutByID(ctx, id)
}

func (i impl) GetServices(ctx echo.Context, params auth.GetServicesParams) error {
	return i.service.Get(ctx, echoserver.GetParams{
		Offset: params.Offset,
		Limit:  params.Limit,
		Sort:   (*repo.Sort)(params.Sort),
	})
}

func (i impl) PostServices(ctx echo.Context) error {
	return i.service.Post(ctx)
}

func (i impl) DeleteServicesId(ctx echo.Context, id auth.Id) error {
	return i.service.DeleteID(ctx, id)
}

func (i impl) GetServicesId(ctx echo.Context, id auth.Id) error {
	return i.service.GetByID(ctx, id)
}

func (i impl) PatchServicesId(ctx echo.Context, id auth.Id) error {
	return i.service.PatchByID(ctx, id)
}

func (i impl) PutServicesId(ctx echo.Context, id auth.Id) error {
	return i.service.PutByID(ctx, id)
}
