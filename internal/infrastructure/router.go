package infrastructure

import (
	"database/sql"
	"muramasa/internal/core"
	productController "muramasa/internal/modules/product/controller"
	stockController "muramasa/internal/modules/stock/controller"
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
	rh.stockAPI(api)
}

func (rh RouterHandler) productAPI(api *gin.RouterGroup) {
	getProductController := productController.NewGetAllProductsController(rh.DB)
	addProductController := productController.NewAddProductController(rh.DB)
	product := api.Group("/product")
	product.GET("/", getProductController.GetAllProducts)
	product.POST("/", addProductController.AddProduct)
}

func (rh RouterHandler) stockAPI(api *gin.RouterGroup) {
	getProductStockByProductIDController := stockController.NewGetProductStockByProductIdController(rh.DB)
	stock := api.Group("/stock")

	stock.GET("/:id", getProductStockByProductIDController.GetProductStockByProductId)
}
