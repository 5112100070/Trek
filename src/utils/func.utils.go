package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RenderingScript(g gin.H, datas ...interface{}) {
	var result string
	for _, data := range datas {
		result = fmt.Sprintf("%s%s", result, data)
	}
	g["script"] = result
}
