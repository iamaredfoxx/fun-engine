package route

import (
	"github.com/labstack/echo/v4"
)

type routePathSingle struct {
	echo *echo.Echo
}

func (r routePathSingle) Route() *echo.Echo {
	return r.echo
}

func (r routePathSingle) Group(group string, f func(rg RoutePathInterface), m ...echo.MiddlewareFunc) {
	g := routePathGroup{echo: r.Route().Group(group, m...)}
	f(g)
}

func (r routePathSingle) Get(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return r.Route().GET(prefix, f, m...)
}

func (r routePathSingle) Post(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return r.Route().POST(prefix, f, m...)
}

func (r routePathSingle) Put(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return r.Route().PUT(prefix, f, m...)
}

func (r routePathSingle) Patch(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return r.Route().PATCH(prefix, f, m...)
}

func (r routePathSingle) Delete(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return r.Route().DELETE(prefix, f, m...)
}
