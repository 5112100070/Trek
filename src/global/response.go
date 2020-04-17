package global

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OKResponse(c *gin.Context, val interface{}) {
	header := gin.H{
		"data": val,
	}
	sendResponse(c, http.StatusOK, header)
}

func ErrorResponse(c *gin.Context, val interface{}) {
	header := gin.H{
		"error": val,
	}
	sendResponse(c, http.StatusOK, header)
}

func CreatedResponse(c *gin.Context, val interface{}) {
	header := gin.H{
		"status_code":    201,
		"server_message": "Success",
		"data":           val,
	}
	sendResponse(c, 201, header)
}

func UnAuthorizeResponse(c *gin.Context, val interface{}) {
	header := gin.H{
		"server_message": "Un Authorize Request",
		"data":           val,
	}
	sendResponse(c, 401, header)
}

func ForbiddenResponse(c *gin.Context, val interface{}) {
	header := gin.H{
		"server_message": "Forbidden",
		"data":           val,
	}
	sendResponse(c, 403, header)
}

func BadRequestResponse(c *gin.Context, val interface{}) {
	header := gin.H{
		"status_code":    400,
		"server_message": "Bad Request",
		"data":           val,
	}
	sendResponse(c, 400, header)
}

func InternalServerErrorResponse(c *gin.Context, val interface{}) {
	header := gin.H{
		"status_code":    500,
		"server_message": "Internal Server Error",
		"data":           val,
	}
	sendResponse(c, 500, header)
}

func sendResponse(c *gin.Context, statusCode int, val gin.H) {
	c.JSON(statusCode, val)
}
