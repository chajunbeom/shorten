package app

import (
	"errors"
	"flag"
	"fmt"
	"reflect"
	"shorten/api"
	"shorten/config"
	"strings"

	"github.com/labstack/echo"
)

type Shortener struct {
	Addr       string
	APIroute   []config.APIroute
	Static     []config.Static
	Echo       *echo.Echo
	APIHandler api.API
}

func (o *Shortener) init() error {
	o.Echo = echo.New()
	v := reflect.ValueOf(o.APIHandler)
	if !v.IsValid() {
		return errors.New("API is nil")
	}

	for _, apiInfo := range o.APIroute {
		h := v.MethodByName(apiInfo.Handler)
		if !h.IsValid() {
			return errors.New("API handler is nil")
		}
		handler, ok := h.Interface().(func(c echo.Context) error)
		if !ok {
			return errors.New("Handler format error")
		}
		o.Echo.Add(strings.ToUpper(apiInfo.Method), apiInfo.Path, handler)
	}

	for _, staticInfo := range o.Static {
		o.Echo.Static(staticInfo.Path, staticInfo.File)
	}
	return nil
}

func (o *Shortener) Start() error {
	return o.Echo.Start(o.Addr)
}

func NewApp() *Shortener {
	configPath := flag.String("c", "config.json", "a config path")
	flag.Parse()

	conf := config.GetInstance(*configPath)
	newShorten := &Shortener{
		APIroute: conf.ServerInfo.APIroute,
		Static:   conf.ServerInfo.Static,
		Addr:     fmt.Sprintf(":%d", conf.ServerInfo.Port),
	}
	if err := newShorten.init(); err != nil {
		panic(err)
	}
	return newShorten
}
