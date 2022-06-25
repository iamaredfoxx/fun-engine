package route

import "github.com/labstack/echo/v4"

type routePathGroup struct {
	echo *echo.Group
}

func (r routePathGroup) Route() *echo.Group {
	return r.echo
}

func (r routePathGroup) Group(group string, f func(rg RoutePathInterface), m ...echo.MiddlewareFunc) {
	g := routePathGroup{echo: r.Route().Group(group, m...)}
	f(g)
}

func (r routePathGroup) Get(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return r.Route().GET(prefix, f, m...)
}

func (r routePathGroup) Post(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return r.Route().POST(prefix, f, m...)
}

func (r routePathGroup) Put(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return r.Route().PUT(prefix, f, m...)
}

func (r routePathGroup) Patch(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return r.Route().PATCH(prefix, f, m...)
}

func (r routePathGroup) Delete(prefix string, f echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return r.Route().DELETE(prefix, f, m...)
}
