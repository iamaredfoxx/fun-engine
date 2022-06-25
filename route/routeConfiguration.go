package route

type RouteConfiguration struct {
	Debug        bool
	HideBanner   bool
	DisableHTTP2 bool
	ReadTimeout  int64
	WriteTimeout int64
}
