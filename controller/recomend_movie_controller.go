package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type searchparams struct {
	adult int `form:"adult"`
	child int `form:"child"`
}

func GetRecomendMovie(c *gin.Context) {
	params, err := validate(c)
	if err != nil {
		log.Println("====== Bind Error ======")
		fmt.Println("error:", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad parameter",
			"detail":  fmt.Sprintf("%s", err),
		})
		return
	}

	log.Println("====== Bind By Query String ======")
	log.Println("adult:", params.adult)
	log.Println("child:", params.child)

	c.JSON(http.StatusOK, gin.H{
		"recomend_movie": "this one",
	})
}

func validate(c *gin.Context) (*searchparams, error) {
	// リクエストパラメーター
	var params searchparams

	// 大人人数
	adultP := c.DefaultQuery("adult", "0")

	if adult, err := customerCheck(adultP, "adult"); err != nil {
		// バリデーションエラー
		return nil, err
	} else {
		params.adult = *adult
	}

	// 子供人数
	childP := c.DefaultQuery("child", "0")

	if child, err := customerCheck(childP, "child"); err != nil {
		// バリデーションエラー
		return nil, err
	} else {
		params.child = *child
	}

	return &params, nil
}

func customerCheck(p string, item string) (*int, error) {
	// 変換後の数値
	var i int
	// エラー
	var err error

	if i, err = strconv.Atoi(p); err != nil {
		return nil, errors.New("parameter must be numeric#" + item)
	}

	if i < 0 {
		return nil, errors.New("parameter must be positive#" + item)
	}

	return &i, nil
}
