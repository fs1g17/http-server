package service

import (
	"net/http"

	"github.com/fs1g17/http-server/instagram"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Getter interface {
	Get() ([]*instagram.Item, error)
}

type Service struct {
	Getter
}

func New(getter Getter) *Service {
	return &Service{
		Getter: getter,
	}
}

func (s *Service) Server() *echo.Echo {
	e := echo.New()
	e.Use(
		middleware.Recover(),
		middleware.Logger(),
		middleware.CORS(),
	)
	e.GET("/", s.items)
	return e
}

func (s *Service) items(c echo.Context) (err error) {
	items, err := s.Get()
	if err != nil {
		c.Logger().Errorf("error getting instagram feed items: %s", err)
		return c.String(http.StatusInternalServerError, "can't fetch instagram items")
	}
	return c.JSON(http.StatusOK, items)
}
