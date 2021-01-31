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
	getProductController := productController.NewGetAllProductsController(rh.Repositories.productRepository)
	addProductController := productController.NewAddProductController(rh.Repositories.productRepository)
	product := api.Group("/product")
	product.GET("/", getProductController.GetAllProducts)
	product.POST("/", addProductController.AddProduct)
}

func (rh RouterHandler) stockAPI(api *gin.RouterGroup) {
	getProductStockByProductIDController := stockController.NewGetProductStockByProductIdController(rh.DB)

	addStockController := stockController.NewAddStockController(
		rh.Repositories.productRepository,
		rh.Repositories.inboundRepository,
		rh.Repositories.stockRepository,
	)
	stock := api.Group("/stock")

	stock.POST("/", addStockController.AddStock)
	stock.GET("/:id", getProductStockByProductIDController.GetProductStockByProductId)
}
