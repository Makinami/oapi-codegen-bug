package app

import "github.com/labstack/echo/v4"

type AppService struct{}

func (a *AppService) GetPeopleBornOnDay(ctx echo.Context, params GetPeopleBornOnDayParams) error {
	return ctx.String(200, params.Day.String())
}
