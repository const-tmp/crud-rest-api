package server

import (
	crud2 "crud-api/crud"
	"crud-api/pkg/crud"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type (
	TransportModel any

	TransformFunc[T1, T2 any] func(*T1) (*T2, error)

	CRUD[T TransportModel, D crud.DBModel] struct {
		crud    *crud.CRUD[D]
		decoder TransformFunc[T, D]
		encoder TransformFunc[D, T]
	}

	// GetParams defines parameters for Get.
	GetParams struct {
		// Offset How many items to skip
		Offset *uint32 `form:"offset,omitempty" json:"offset,omitempty"`

		// Limit How many items to return at one time
		Limit *uint32 `form:"limit,omitempty" json:"limit,omitempty"`

		// Sort order
		Sort *map[string]string `form:"sort,omitempty" json:"sort,omitempty"`
	}
)

func New[T TransportModel, D crud.DBModel](db *gorm.DB, decoder TransformFunc[T, D], encoder TransformFunc[D, T]) *CRUD[T, D] {
	return &CRUD[T, D]{
		crud:    crud.New[D](db),
		encoder: encoder,
		decoder: decoder,
	}
}

func (c CRUD[T, D]) Get(ctx echo.Context, params GetParams) error {
	rows, err := c.crud.Get(ctx.Request().Context(), params.Offset, params.Limit, params.Sort)
	if err != nil {
		return err
	}

	result := make([]*T, len(rows))
	for _, row := range rows {
		if res, err := c.encoder(row); err != nil {
			return err
		} else {
			result = append(result, res)
		}
	}

	return ctx.JSON(http.StatusOK, rows)
}

func (c CRUD[T, D]) Post(ctx echo.Context) error {
	var v *T
	if err := ctx.Bind(&v); err != nil {
		return err
	}

	row, err := c.decoder(v)
	if err != nil {
		return err
	}

	row, err = c.crud.Create(ctx.Request().Context(), *row)
	if err != nil {
		return err
	}

	if res, err := c.encoder(row); err != nil {
		return err
	} else {
		return ctx.JSON(http.StatusCreated, res)
	}
}

func (c CRUD[T, D]) DeleteID(ctx echo.Context, id crud2.Id) error {
	if err := c.crud.DeleteByID(ctx.Request().Context(), uint(id)); err != nil {
		return err
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (c CRUD[T, D]) GetByID(ctx echo.Context, id crud2.Id) error {
	v, err := c.crud.GetByID(ctx.Request().Context(), uint(id))
	if err != nil {
		return err
	}

	if res, err := c.encoder(v); err != nil {
		return err
	} else {
		return ctx.JSON(http.StatusOK, res)
	}
}

func (c CRUD[T, D]) PatchByID(ctx echo.Context, id crud2.Id) error {
	var m = make(echo.Map)
	if err := ctx.Bind(&m); err != nil {
		return err
	}
	if err := c.crud.Update(ctx.Request().Context(), uint(id), m); err != nil {
		return err
	}
	return ctx.NoContent(http.StatusOK)
}

func (c CRUD[T, D]) PutByID(ctx echo.Context, id crud2.Id) error {
	var v T
	if err := ctx.Bind(&v); err != nil {
		return err
	}

	row, err := c.decoder(&v)
	if err != nil {
		return err
	}

	if err := c.crud.Replace(ctx.Request().Context(), uint(id), *row); err != nil {
		return err
	}
	return ctx.NoContent(http.StatusOK)
}
