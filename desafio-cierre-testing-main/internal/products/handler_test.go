package products

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func createServer(mockRepository mocks.MockRepository) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	service := products.NewService(&mockRepository)
	handler := NewHandler(service)

	r := gin.Default()

	pr := r.Group("/products")
	pr.GET("/", handler.GetProducts)

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}
