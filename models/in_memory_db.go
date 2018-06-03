package models

import (
	"errors"
	"sync"
)

type data *interface{}
type memoryDB map[interface{}]data

var instanceMemoryDB *memoryDB
var onceMemoryDB sync.Once

func GetMemoryDB() *memoryDB {
	onceMemoryDB.Do(func() {
		instanceMemoryDB = &memoryDB{}
	})
	return instanceMemoryDB
}

func (o *memoryDB) get(key interface{}) data {
	return (*o)[key]
}

func (o *memoryDB) create(key interface{}, d data) error {
	if _, ok := (*o)[key]; ok {
		return errors.New("data exit")
	}
	(*o)[key] = d
	return nil
}

func (o *memoryDB) update(key interface{}, d data) error {
	if _, ok := (*o)[key]; !ok {
		return errors.New("do not exit")
	}
	(*o)[key] = d
	return nil
}

func (o *memoryDB) remove(key interface{}) error {
	if _, ok := (*o)[key]; !ok {
		return errors.New("do not exit")
	}
	delete(*o, key)
	return nil
}

// Interface
func (o *memoryDB) GetOriginURL(key string) string {
	val := o.get(key)
	if val == nil {
		return ""
	}
	v := (*val).(*ShortenURLdata)
	return v.Origin
}

func (o *memoryDB) GetShortenURL(key string) string {
	val := o.get(key)
	if val == nil {
		return ""
	}
	v := (*val).(*ShortenURLdata)
	return v.Shorten
}

func (o *memoryDB) SetURL(key string, val interface{}) error {
	return o.create(key, &val)
}

func (o *memoryDB) DeletURL(key string) error {
	return o.remove(key)
}
