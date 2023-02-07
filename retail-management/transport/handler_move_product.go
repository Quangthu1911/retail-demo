package transport

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	retailbiz "retail-demo/retail-management/business"
	"retail-demo/retail-management/model/dto"
	retailStorage "retail-demo/retail-management/storage"
	"strings"
)

func HandleMoveProducts(db *mongo.Database) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto dto.MoveProductDto

		if err := ctx.ShouldBind(&dto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if dto.OutputInventory == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument"})
			return
		}
		if len(dto.InputInventory.Products) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "please choose product"})
			return
		}

		dto.OutputInventory = strings.TrimSpace(dto.OutputInventory)

		storage := retailStorage.NewMongoDbStorage(db)
		biz := retailbiz.MoveProductBiz(storage)

		err := biz.MoveProduct(ctx.Request.Context(), dto)
		if err != nil {
			if errors.Is(err, errors.New("user not found")) {
				ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"result": "Success"})
	}
}
