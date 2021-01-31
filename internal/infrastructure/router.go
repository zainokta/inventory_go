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
	repositories := initRepositories(db)
	return RouterHandler{
		Logger:       logger,
		DB:           db,
		Repositories: repositories,
	}
}

func (rh RouterHandler) SetRoutes(r *gin.Engine) {

	api := r.Group("/api")
	api.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	rh.productAPI(api)
	rh.stockAPI(api)
	rh.orderAPI(api)
}

func (rh RouterHandler) productAPI(api *gin.RouterGroup) {
	product := api.Group("/product")

	productController := productController.NewProductController(rh.Repositories.productRepository)

	product.GET("/", productController.GetAllProducts)
	product.POST("/", productController.AddProduct)
}

func (rh RouterHandler) stockAPI(api *gin.RouterGroup) {
	stock := api.Group("/stock")

	stockController := stockController.NewStockController(rh.Repositories.productRepository, rh.Repositories.inboundRepository, rh.Repositories.stockRepository)

	stock.POST("/", stockController.AddStock)
	stock.GET("/:id", stockController.GetProductStockByProductId)
}

func (rh RouterHandler) orderAPI(api *gin.RouterGroup) {
	order := api.Group("/order")

	order.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "order")
	})
}
