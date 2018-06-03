package models

import (
	"shorten/config"
	"sync"
)

type DatatBase interface {
	GetOriginURL(key string) string
	GetShortenURL(key string) string
	SetURL(key string, val interface{}) error
	DeletURL(key string) error
}

var db DatatBase
var once sync.Once

func getInstance() DatatBase {
	once.Do(func() {
		conf := config.GetInstance()
		switch conf.DBtype {
		case "memory":
			db = GetMemoryDB()
		default:
			db = GetMemoryDB()
		}
	})
	return db
}

func GetOriginURL(shorten string) string {
	d := getInstance()
	return d.GetOriginURL(shorten)
}

func GetShortenURL(origin string) string {
	d := getInstance()
	return d.GetShortenURL(origin)
}

func SetURL(key string, val interface{}) error {
	d := getInstance()
	return d.SetURL(key, val)
}

func DeletURL(key string) error {
	d := getInstance()
	return d.DeletURL(key)
}

// Schema
type ShortenURLdata struct {
	Origin  string `json:"origin"`
	Shorten string `json:"shorten_url"`
}
