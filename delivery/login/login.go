package login

import "mncTestApp/usecase"
import "mncTestApp/delivery/apprequest"
import "github.com/gin-gonic/gin"
import "mncTestApp/delivery/commonresp"
import "net/http"

type loginApi struct {
	usecase usecase.AuthenticationUseCase
}

func (api *loginApi) Login() gin.HandlerFunc {
	return func (c *gin.Context){
		var payload apprequest.LoginRequest
		err := c.ShouldBindJSON(&payload)
		if err != nil {
			commonresp.NewJsonResponse(c).SendError(commonresp.NewErrorMessage(http.StatusInternalServerError, "userName and Password is required"))
		}
		errLogin := api.usecase.Login(payload.Username, payload.Password)

		if errLogin != nil {
			commonresp.NewJsonResponse(c).SendError(commonresp.NewErrorMessage(http.StatusBadRequest, "username and password does'nt match"))
		}

		commonresp.NewJsonResponse(c).SendData(commonresp.NewResponseMessage(http.StatusOK, "Login Success", payload))


	}
}


