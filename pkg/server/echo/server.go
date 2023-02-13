package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/nullc4t/crud-rest-api/pkg/common"
	"github.com/nullc4t/crud-rest-api/pkg/repo"
	"gorm.io/gorm"
	"net/http"
)

type (
	Server[T, R any] struct {
		repo    *repo.CRUD[R]
		decoder common.TransformFunc[T, R]
		encoder common.TransformFunc[R, T]
	}

	// GetParams defines parameters for Get.
	GetParams struct {
		// Offset How many items to skip
		Offset *uint32 `form:"offset,omitempty" json:"offset,omitempty"`

		// Limit How many items to return at one time
		Limit *uint32 `form:"limit,omitempty" json:"limit,omitempty"`

		// Sort order
		Sort *repo.Sort `form:"sort,omitempty" json:"sort,omitempty"`
	}
)

func New[T, D any](db *gorm.DB, decoder common.TransformFunc[T, D], encoder common.TransformFunc[D, T]) *Server[T, D] {
	return &Server[T, D]{
		repo:    repo.New[D](db),
		encoder: encoder,
		decoder: decoder,
	}
}

func (c Server[T, D]) Get(ctx echo.Context, params GetParams) error {
	rows, err := c.repo.Get(ctx.Request().Context(), params.Offset, params.Limit, params.Sort)
	if err != nil {
		return err
	}

	result := make([]*T, 0, len(rows))
	for _, row := range rows {
		if res, err := c.encoder(row); err != nil {
			return err
		} else {
			result = append(result, res)
		}
	}

	return ctx.JSON(http.StatusOK, result)
}

func (c Server[T, D]) Post(ctx echo.Context) error {
	var v *T
	if err := ctx.Bind(&v); err != nil {
		return err
	}

	row, err := c.decoder(v)
	if err != nil {
		return err
	}

	row, err = c.repo.Create(ctx.Request().Context(), *row)
	if err != nil {
		return err
	}

	if res, err := c.encoder(row); err != nil {
		return err
	} else {
		return ctx.JSON(http.StatusCreated, res)
	}
}

func (c Server[T, D]) DeleteID(ctx echo.Context, id uint64) error {
	if err := c.repo.DeleteByID(ctx.Request().Context(), uint(id)); err != nil {
		return err
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (c Server[T, D]) GetByID(ctx echo.Context, id uint64) error {
	v, err := c.repo.GetByID(ctx.Request().Context(), uint(id))
	if err != nil {
		return err
	}

	if res, err := c.encoder(v); err != nil {
		return err
	} else {
		return ctx.JSON(http.StatusOK, res)
	}
}

func (c Server[T, D]) PatchByID(ctx echo.Context, id uint64) error {
	var m = make(echo.Map)
	if err := ctx.Bind(&m); err != nil {
		return err
	}
	if err := c.repo.Update(ctx.Request().Context(), uint(id), m); err != nil {
		return err
	}
	return ctx.NoContent(http.StatusOK)
}

func (c Server[T, D]) PutByID(ctx echo.Context, id uint64) error {
	var v T
	if err := ctx.Bind(&v); err != nil {
		return err
	}

	row, err := c.decoder(&v)
	if err != nil {
		return err
	}

	if err = c.repo.Replace(ctx.Request().Context(), uint(id), *row); err != nil {
		return err
	}
	return ctx.NoContent(http.StatusOK)
}
