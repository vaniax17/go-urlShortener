package handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/vaniax17/go-urlShortener/internal/domain"
)

type UserHandler struct {
	userService domain.UserService
}

func NewUserHandler(userService domain.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

type userRequest struct {
	Username              string `json:"username"`
	Base64EncodedPassword string `json:"password"`
}

type userResponse struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}

func (u *UserHandler) CreateUser(c *echo.Context) error {
	var req userRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid json body")
	}

	ctx := c.Request().Context()
	user, tokenPair, err := u.userService.Create(ctx, req.Username, req.Base64EncodedPassword)
	if err != nil {
		return handleError(err)
	}

	setCookieAuthCookie(c, tokenPair)
	return c.JSON(http.StatusCreated, userResponse{
		Id:       user.Id,
		Username: user.Username,
	})

}

func (u *UserHandler) LoginUser(c *echo.Context) error {
	var req userRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid json body")
	}

	ctx := c.Request().Context()
	user, tokenPair, err := u.userService.Login(ctx, req.Username, req.Base64EncodedPassword)
	if err != nil {
		return handleError(err)
	}

	setCookieAuthCookie(c, tokenPair)
	return c.JSON(http.StatusCreated, user)
}

func setCookieAuthCookie(c *echo.Context, tokenPair *domain.TokenPair) {
	c.SetCookie(&http.Cookie{
		Name:     "access_token",
		Value:    tokenPair.AccessToken,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		MaxAge:   86400, // 86400 = 24hours
	})

	c.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    tokenPair.RefreshToken,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		MaxAge:   86400 * 7, // 86400 = 24hours
	})
}
