package test

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"strings"
	"testing"
	"fmt"
)

func TestHoge(t *testing.T) {

	// ポストデータ
	bodyReader := strings.NewReader(`{"password": "test"}`)
	// リクエスト生成
	req := httptest.NewRequest("POST", "/", bodyReader)
	// Content-Type 設定
	req.Header.Set("Content-Type", "application/json")

	// Contextセット
	var context *gin.Context
	context = &gin.Context{Request: req}
	context.Set("JWT_PAYLOAD",jwt.MapClaims{
		"id":float64(10),
	})

	huga(context)
}

func huga(c *gin.Context) {
	userIdUint :=  uint64(jwt.ExtractClaims(c)["id"].(float64))
	type SSS struct{
		Password string	`json:"password"`
	}
	var sss SSS
	c.Bind(&sss)
	fmt.Println(sss)
	fmt.Println(userIdUint)
}
