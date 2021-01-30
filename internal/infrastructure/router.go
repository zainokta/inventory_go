package infrastructure

import (
	"database/sql"
	"muramasa/internal/core"
	"muramasa/internal/modules/product/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RouterHandler struct {
	Logger core.ILogger
	DB     *sql.DB
}

func NewRouterWithLogger(logger core.ILogger, db *sql.DB) RouterHandler {
	return RouterHandler{
		Logger: logger,
		DB:     db,
	}
}

func (rh RouterHandler) SetRoutes(r *gin.Engine) {

	api := r.Group("/api")
	api.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	rh.productAPI(api)
}

func (rh RouterHandler) productAPI(api *gin.RouterGroup) {
	getProductController := controller.NewGetAllProductsController(rh.DB)
	addProduct := controller.NewAddProductController(rh.DB)
	product := api.Group("/product")
	product.GET("/", getProductController.GetAllProducts)
	product.POST("/", addProduct.AddProduct)
}
