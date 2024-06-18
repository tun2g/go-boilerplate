package auth


import(
	httpContext "fist-app/src/shared/http-context"
	_ "fist-app/src/apis/dto/auth"
	"fist-app/src/shared/exception"
)


func initException(_ exception.HttpError){}

// @Summary Login
// @Description User login
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param   loginReq  body  _.LoginReqDto  true  "Login request"
// @Success 200 {object} _.AuthResDto
// @Failure 422 {object} exception.HttpError
// @Failure 400 {object} exception.HttpError
// @Router /auth/sign-in [post]
func login(ctx *httpContext.CustomContext) {}

// @Summary Register
// @Description User Register
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param   registerReq  body  _.RegisterReqDto  true  "Register request"
// @Success 201 {object} _.AuthResDto
// @Failure 422 {object} exception.HttpError
// @Failure 400 {object} exception.HttpError
// @Router /auth/sign-up [post]
func register(ctx *httpContext.CustomContext){}