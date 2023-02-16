package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/nullc4t/crud-rest-api/pkg/common"
	"github.com/nullc4t/crud-rest-api/pkg/repo"
	gormrepo "github.com/nullc4t/crud-rest-api/pkg/repo/gorm"
	"gorm.io/gorm"
	"net/http"
)

type (
	Server[T, R any] struct {
		repo    repo.Interface[R]
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

func New[T, R any](db *gorm.DB, decoder common.TransformFunc[T, R], encoder common.TransformFunc[R, T]) *Server[T, R] {
	return &Server[T, R]{
		repo:    gormrepo.New[R](db),
		encoder: encoder,
		decoder: decoder,
	}
}

func (s Server[T, R]) Get(ctx echo.Context, params GetParams) error {
	rows, err := s.repo.Get(ctx.Request().Context(), params.Offset, params.Limit, params.Sort)
	if err != nil {
		return err
	}

	result := make([]*T, 0, len(rows))
	for _, row := range rows {
		if res, err := s.encoder(row); err != nil {
			return err
		} else {
			result = append(result, res)
		}
	}

	return ctx.JSON(http.StatusOK, result)
}

func (s Server[T, R]) Post(ctx echo.Context) error {
	var v *T
	if err := ctx.Bind(&v); err != nil {
		return err
	}

	row, err := s.decoder(v)
	if err != nil {
		return err
	}

	row, err = s.repo.Create(ctx.Request().Context(), *row)
	if err != nil {
		return err
	}

	if res, err := s.encoder(row); err != nil {
		return err
	} else {
		return ctx.JSON(http.StatusCreated, res)
	}
}

func (s Server[T, R]) DeleteID(ctx echo.Context, id uint64) error {
	if err := s.repo.DeleteByID(ctx.Request().Context(), uint(id)); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (s Server[T, R]) GetByID(ctx echo.Context, id uint64) error {
	v, err := s.repo.GetByID(ctx.Request().Context(), uint(id))
	if err != nil {
		return err
	}

	if res, err := s.encoder(v); err != nil {
		return err
	} else {
		return ctx.JSON(http.StatusOK, res)
	}
}

func (s Server[T, R]) PatchByID(ctx echo.Context, id uint64) error {
	var m = make(echo.Map)
	if err := ctx.Bind(&m); err != nil {
		return err
	}

	if err := s.repo.Update(ctx.Request().Context(), uint(id), m); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}

func (s Server[T, R]) PutByID(ctx echo.Context, id uint64) error {
	var v T
	if err := ctx.Bind(&v); err != nil {
		return err
	}

	row, err := s.decoder(&v)
	if err != nil {
		return err
	}

	if err = s.repo.Replace(ctx.Request().Context(), uint(id), *row); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}
