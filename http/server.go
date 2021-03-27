package http

import (
	"example/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WebServer struct {
	Products *model.Products
	Port     string
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	e := echo.New()
	e.GET("/products", w.getAll)
	e.POST("/products", w.createProduct)
	e.Start(w.Port)
}

func (w WebServer) getAll(c echo.Context) error {
	return c.JSON(http.StatusOK, w.Products)
}

func (w WebServer) createProduct(c echo.Context) error {
	product := model.NewProduct()
	if err := c.Bind(product); err != nil {
		return err
	}
	w.Products.Add(product)
	return c.JSON(http.StatusCreated, product)
}
