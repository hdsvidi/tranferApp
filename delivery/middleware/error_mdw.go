package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"mncTestApp/delivery/commonresp"
	"mncTestApp/logger"
	"net/http"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedError := c.Errors.Last()
		if detectedError == nil {
			return
		}
		e := detectedError.Error()
		errResp := commonresp.ErrorMessage{}
		err := json.Unmarshal([]byte(e), &errResp)
		if err != nil {
			errResp.HttpCode = http.StatusInternalServerError
			errResp.ErrorDescription = commonresp.ErrorDescription{
				Status:       "Error",
				Message:  "Convert Json Failed",
			}
			logger.Log.Error().Err(err).Msg(errResp.Message)
		} else {
			logger.Log.Error().Err(fmt.Errorf("%d", errResp.HttpCode)).Str("something we're wrong", errResp.Message)
		}

		commonresp.NewJsonResponse(c).SendError(errResp)
	}
}
