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

func HandleImportProduct(db *mongo.Database) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto dto.ImportProductDto

		if err := ctx.ShouldBind(&dto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if dto.InventoryId == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument branchName"})
			return
		}
		if len(dto.Products) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "please choose product"})
			return
		}

		dto.InventoryId = strings.TrimSpace(dto.InventoryId)

		storage := retailStorage.NewMongoDbStorage(db)
		biz := retailbiz.ImportProductBiz(storage)

		err := biz.ImportProduct(ctx.Request.Context(), dto)
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
