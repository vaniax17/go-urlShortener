package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/vaniax17/go-urlShortener/internal/domain"
)

func handleError(err error) error {
	switch {
	case errors.Is(err, domain.ErrUserAlreadyExists):
		return echo.NewHTTPError(http.StatusConflict, domain.ErrUserAlreadyExists.Error())
	case errors.Is(err, domain.ErrDecodingBase64):
		return echo.NewHTTPError(http.StatusBadRequest, domain.ErrDecodingBase64.Error())
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

}
