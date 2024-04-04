package warehouse

import (
	"e-commerce/internal/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetWarehouse(ctx *gin.Context) {
	id := ctx.Param("id")

	warehouse, err := h.uc.GetWarehouse(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":         warehouse.ID,
			"city":       warehouse.City,
			"status":     warehouse.Status,
			"stocks":     warehouse.Stocks,
			"created_at": warehouse.CreatedAt,
			"updated_at": warehouse.UpdatedAt,
		},
	})
}

func (h Handler) CreateWarehouse(ctx *gin.Context) {
	var input entity.CreateWarehouse

	err := ctx.BindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	warehouse, err := h.uc.CreateWarehouse(ctx.Request.Context(), input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":         warehouse.ID,
			"city":       warehouse.City,
			"status":     warehouse.Status,
			"created_at": warehouse.CreatedAt,
			"updated_at": warehouse.UpdatedAt,
		},
	})
}

func (h Handler) DeleteWarehouse(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.uc.DeleteWarehouse(ctx.Request.Context(), id)
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

func (h Handler) UpdateStatusWarehouse(ctx *gin.Context) {
	var input entity.UpdateWarehouse

	err := ctx.BindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 32, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	input.ID = uint(id)

	warehouse, err := h.uc.UpdateStatusWarehouse(ctx.Request.Context(), input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":         warehouse.ID,
			"city":       warehouse.City,
			"status":     warehouse.Status,
			"created_at": warehouse.CreatedAt,
			"updated_at": warehouse.UpdatedAt,
		},
	})
}
