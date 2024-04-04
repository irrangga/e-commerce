package order

import (
	"e-commerce/internal/entity"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) CreateOrder(ctx *gin.Context) {
	var input entity.CreateOrder

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

	order, err := h.uc.CreateOrder(ctx.Request.Context(), input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":             order.ID,
			"product_orders": order.ProductOrders,
			"created_at":     order.CreatedAt,
			"updated_at":     order.UpdatedAt,
		},
	})
}

func (h Handler) DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.uc.DeleteOrder(ctx.Request.Context(), id)
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

func (h Handler) CheckoutOrder(ctx *gin.Context) {
	id := ctx.Param("id")

	order, err := h.uc.CheckoutOrder(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":             order.ID,
			"product_orders": order.ProductOrders,
			"created_at":     order.CreatedAt,
			"updated_at":     order.UpdatedAt,
		},
	})
}
