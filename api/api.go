package api

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/labstack/echo"

	"shorten/models"
	"shorten/utils"
)

const (
	urlCheck = "^(http:\\/\\/www\\.|https:\\/\\/www\\.|http:\\/\\/|https:\\/\\/)?[a-z0-9]+([\\-\\.]{1}[a-z0-9]+)*\\.[a-z]{2,5}(:[0-9]{1,5})?(\\/.*)?"
)

type API struct {
}

func (o API) Redirecting(c echo.Context) error {
	shortURL := c.Param("shortURL")
	fmt.Println(shortURL)
	if shortURL == "" {
		resp := Response{}
		resp.Error("short url invalid", nil)
		return c.JSON(http.StatusOK, resp)
	}
	redirectURL := models.GetOriginURL(shortURL)
	if redirectURL == "" {
		fmt.Println("not found")
		return c.NoContent(404)
	}

	fmt.Println("redirectingURL", redirectURL)
	if !strings.HasPrefix(redirectURL, "https://") && !strings.HasPrefix(redirectURL, "http://") {
		redirectURL = "http://" + redirectURL
	}

	return c.Redirect(http.StatusPermanentRedirect, redirectURL)
}

func (o API) ConvertURL(c echo.Context) error {
	origin := c.FormValue("origin_url")

	resp := Response{}
	if origin == "" {
		resp.Error("origin_url field null", nil)
		return c.JSON(http.StatusOK, resp)
	}

	rg := regexp.MustCompile(urlCheck)
	if !rg.Match([]byte(origin)) {
		resp.Error("Invalid url", nil)
		return c.JSON(http.StatusOK, resp)
	}
	data := &models.ShortenURLdata{
		Origin: origin,
	}

	fmt.Println("original URL", data.Origin)
	if short := models.GetShortenURL(origin); short != "" {
		data.Shorten = short
		resp.OK(data)
		return c.JSON(http.StatusOK, resp)
	}

	coll := uint16(0)
	hash := utils.CRC32([]byte(origin))

	shortRaw, _ := utils.ConvertIntToByte(utils.INT16, uint64(coll))
	hashByte, _ := utils.ConvertIntToByte(utils.INT32, uint64(hash))
	short := utils.Base64Encode(append(shortRaw, hashByte...))

	for models.GetOriginURL(string(short)) != "" {
		coll++
		if coll == uint16(0xFF) {
			resp.Error("hash chain full", nil)
			return c.JSON(http.StatusOK, resp)
		}
		temp, _ := utils.ConvertIntToByte(utils.INT16, uint64(coll))
		copy(shortRaw[0:2], temp[0:2])
		short = utils.Base64Encode(shortRaw)
	}

	data.Shorten = strings.Trim(string(short), string(0))

	if err := models.SetURL(data.Origin, data); err != nil {
		resp.Error("db error", nil)
		return c.JSON(http.StatusOK, resp)
	}
	if err := models.SetURL(data.Shorten, data); err != nil {
		go models.DeletURL(data.Origin)
		resp.Error("db error", nil)
		return c.JSON(http.StatusOK, resp)
	}

	resp.OK(data)
	return c.JSON(http.StatusOK, resp)
}
