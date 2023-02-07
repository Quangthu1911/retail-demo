package transport

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	retailbiz "retail-demo/retail-management/business"
	retailStorage "retail-demo/retail-management/storage"
)

func HandleGetInfo(db *mongo.Database) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		storage := retailStorage.NewMongoDbStorage(db)
		biz := retailbiz.GetInfoBiz(storage)
		err := biz.GetInfo(ctx.Request.Context())
		ctx.JSON(http.StatusOK, gin.H{"result": err})
	}
}
