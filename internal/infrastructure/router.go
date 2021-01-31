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
	Logger       core.ILogger
	DB           *sql.DB
	Repositories *Repositories
}

func NewRouterWithLogger(logger core.ILogger, db *sql.DB) RouterHandler {
	return RouterHandler{
		Logger:       logger,
		DB:           db,
		Repositories: initRepositories(db),
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
	product := api.Group("/product")

	product.GET("/", productController.NewGetAllProductsController(rh.Repositories.productRepository).GetAllProducts)
	product.POST("/", productController.NewAddProductController(rh.Repositories.productRepository).AddProduct)
}

func (rh RouterHandler) stockAPI(api *gin.RouterGroup) {
	stock := api.Group("/stock")

	stock.POST("/", stockController.NewAddStockController(
		rh.Repositories.productRepository,
		rh.Repositories.inboundRepository,
		rh.Repositories.stockRepository,
	).AddStock)
	stock.GET("/:id", stockController.NewGetProductStockByProductIdController(rh.Repositories.stockRepository).GetProductStockByProductId)
}
