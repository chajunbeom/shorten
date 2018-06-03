package models_test

import (
	"fmt"
	"shorten/config"
	"shorten/models"
	"testing"
)

var conf = config.GetInstance("../etc/conf/config.json")

const TestCase = 100

func TestDataGetSet(t *testing.T) {
	db := models.GetMemoryDB()
	if db == nil {
		t.Error("DB instance null")
	}

	TestKeyValueFair := make([]models.ShortenURLdata, 0)
	for i := 0; i < TestCase; i++ {
		data := models.ShortenURLdata{
			Origin:  fmt.Sprintf("key%d", i),
			Shorten: fmt.Sprintf("value%d", i),
		}
		TestKeyValueFair = append(TestKeyValueFair, data)
		if err := db.SetURL(data.Origin, &data); err != nil {
			t.Error("database error")
		}
		if err := db.SetURL(data.Shorten, &data); err != nil {
			t.Error("database error")
		}
	}

	for _, v := range TestKeyValueFair {
		if db.GetOriginURL(v.Shorten) != v.Origin {
			t.Fail()
		}
		if db.GetShortenURL(v.Origin) != v.Shorten {
			t.Fail()
		}
	}
	fmt.Println(t.Name() + ":OK")
}
