package route

import "github.com/labstack/echo/v4"

type RoutePathInterface interface {
	Group(group string, f func(rg RoutePathInterface), m ...echo.MiddlewareFunc)
	Get(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	Post(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	Put(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	Patch(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	Delete(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

type RouteInterface interface {
	Create(RoutePathInterface)
}
