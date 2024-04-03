package product

import (
	"e-commerce/internal/entity"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetListProducts(ctx *gin.Context) {
	products, err := h.uc.GetListProducts(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var data []gin.H
	for _, product := range products {
		data = append(data, gin.H{
			"id":                 product.ID,
			"name":               product.Name,
			"price":              product.Price,
			"stock_availability": product.StockAvailability,
			"created_at":         product.CreatedAt,
			"updated_at":         product.UpdatedAt,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (h Handler) GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	product, err := h.uc.GetProduct(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":                 product.ID,
			"name":               product.Name,
			"price":              product.Price,
			"stock_availability": product.StockAvailability,
			"created_at":         product.CreatedAt,
			"updated_at":         product.UpdatedAt,
		},
	})
}

func (h Handler) CreateProduct(ctx *gin.Context) {
	var input entity.CreateProduct

	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	err = json.Unmarshal(body, &input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	product, err := h.uc.CreateProduct(ctx.Request.Context(), input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":         product.ID,
			"name":       product.Name,
			"price":      product.Price,
			"stocks":     product.Stocks,
			"created_at": product.CreatedAt,
			"updated_at": product.UpdatedAt,
		},
	})
}

func (h Handler) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.uc.DeleteProduct(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": nil,
	})
}
