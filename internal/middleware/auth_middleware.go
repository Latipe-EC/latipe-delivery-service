package middleware

import (
	"delivery-service/internal/adapter/userserv"
	"delivery-service/internal/adapter/userserv/dto"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	userServ *userserv.UserService
}

func NewAuthMiddleware(service *userserv.UserService) *AuthMiddleware {
	return &AuthMiddleware{userServ: service}
}

func (auth AuthMiddleware) RequiredRoles(roles []string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		bearToken := ctx.Get("Authorization")
		if bearToken == "" {
			return ctx.Status(http.StatusUnauthorized).SendString("Unauthenticated")
		}

		bearToken = strings.Split(bearToken, " ")[1]
		req := dto.AuthorizationRequest{}
		req.Token = bearToken

		resp, err := auth.userServ.Authorization(ctx.Context(), &req)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).SendString("Internal Server Error")
		}

		ctx.Locals(BEARER_TOKEN, bearToken)
		ctx.Locals(USER_ID, resp.Id)

		for _, i := range roles {
			if i == resp.Role {
				return ctx.Next()
			}
		}
		return ctx.Status(http.StatusForbidden).SendString("Permission Denied")
	}
}
