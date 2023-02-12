package utils

import (
	"github.com/gin-gonic/gin"
)

// type HttpMethod string

// const (
// 	GET    HttpMethod = "GET"
// 	POST   HttpMethod = "POST"
// 	PUT    HttpMethod = "PUT"
// 	DELETE HttpMethod = "DELETE"
// )

// func CheckMethod(httpMethod string, chkMethod HttpMethod) bool {
// 	return httpMethod == string(chkMethod)
// }

// func GetUrlId(r *http.Request) (int, error) {
// 	id, err := strconv.Atoi(r.URL.Query().Get("id"))
// 	return id, err
// }

// func GetUrlParmId(r *http.Request) (int, error) {
// 	path := r.URL.Path
// 	isStr := strings.Split(path, "/")
// 	id, err := strconv.Atoi(isStr[len(isStr)-1])
// 	return id, err
// }

func CustomRepsonseWriter(c *gin.Context, status int, data any, message string) {
	c.JSON(status, gin.H{
		"data":    data,
		"message": message,
		"status":  status,
	})
}
