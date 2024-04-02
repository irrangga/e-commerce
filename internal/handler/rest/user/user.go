package user

import (
	"e-commerce/internal/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := h.uc.GetUser(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":           user.ID,
			"name":         user.Name,
			"email":        user.Email,
			"phone_number": user.PhoneNumber,
			"created_at":   user.CreatedAt,
			"updated_at":   user.UpdatedAt,
		},
	})
}

func (h Handler) CreateUser(ctx *gin.Context) {
	var input entity.CreateUser

	err := ctx.BindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.uc.CreateUser(ctx.Request.Context(), input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":           user.ID,
			"name":         user.Name,
			"email":        user.Email,
			"phone_number": user.PhoneNumber,
			"created_at":   user.CreatedAt,
			"updated_at":   user.UpdatedAt,
		},
	})
}

func (h Handler) UpdateUser(ctx *gin.Context) {
	var input entity.UpdateUser

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

	user, err := h.uc.UpdateUser(ctx.Request.Context(), input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":           user.ID,
			"name":         user.Name,
			"email":        user.Email,
			"phone_number": user.PhoneNumber,
			"created_at":   user.CreatedAt,
			"updated_at":   user.UpdatedAt,
		},
	})
}

func (h Handler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := h.uc.DeleteUser(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":           user.ID,
			"name":         user.Name,
			"email":        user.Email,
			"phone_number": user.PhoneNumber,
			"created_at":   user.CreatedAt,
			"updated_at":   user.UpdatedAt,
		},
	})
}
