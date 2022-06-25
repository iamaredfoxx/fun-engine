package route

import (
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

var routeInstance *route

func Route() *route {
	if routeInstance == nil {
		routeInstance = &route{
			echo:   echo.New(),
			config: &RouteConfiguration{},
		}
	}
	return routeInstance
}

func Register(config RouteConfiguration, paths []RouteInterface) {
	Route().config = &config
	Route().prep()
	for _, path := range paths {
		singleRoutePathInstance := &routePathSingle{
			echo: Route().echo,
		}
		path.Create(singleRoutePathInstance)
	}
}

func Bind(port string) {
	Route().route().Start(strings.Join([]string{":", port}, ""))
}

type route struct {
	echo   *echo.Echo
	config *RouteConfiguration
}

func (r *route) route() *echo.Echo {
	return r.echo
}

func (r *route) prep() {
	r.route().HideBanner = r.config.HideBanner
	r.route().DisableHTTP2 = r.config.DisableHTTP2
	r.route().Debug = r.config.Debug
	r.route().Server.ReadTimeout = time.Duration(r.config.ReadTimeout)
	r.route().Server.WriteTimeout = time.Duration(r.config.WriteTimeout)
}
