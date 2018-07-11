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

func BuildMessageForRequest(productID int64, productName string, typeDuration string, duration int64, total int64, startDate string, email string, projectAddress string) string {
	var result string
	result = fmt.Sprintf("Product ID: %d \n", productID)
	result = fmt.Sprintf("%sProduct NAME: %s \n", result, productName)
	result = fmt.Sprintf("%sTipe Durasi: %s \n", result, typeDuration)
	result = fmt.Sprintf("%sDurasi: %d \n", result, duration)
	result = fmt.Sprintf("%sTotal: %d \n", result, total)
	result = fmt.Sprintf("%sTanggal Mulai: %s \n", result, startDate)
	result = fmt.Sprintf("%sEmail: %s \n", result, email)
	result = fmt.Sprintf("%sAlamat Proyek: %s \n", result, projectAddress)

	return result
}
