package ops

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

// Handling handles an error by setting a message and a response status code
func Handling(ctx *gin.Context, err error) {
	var e *MyError

	if !errors.As(err, &e) {
		Handling(ctx, Err(err))
		return
	}
	log.Printf("%s error: %s", e.Location, e.Error())
	e.Message = err.Error()
	ctx.JSON(int(e.HttpStatusCode), gin.H{
		"message": e.Error(),
	})
	ctx.Set("error", err.Error())
	ctx.Abort()
}
