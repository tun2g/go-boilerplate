package auth


import(
	httpContext "fist-app/src/shared/http-context"
	_ "fist-app/src/apis/dto/auth"

)

// @Tags         Auth
// @Summary      Login
// @Description  Login
// @Accept       json
// @Produce      json
// @Success      200  {object}   _.AuthResDto
// @Router       /sign-in [post]
func login(ctx *httpContext.CustomContext) {}
