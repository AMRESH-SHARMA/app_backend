package agora

import (
	"app_backend/internal/response"
	"os"

	rtctokenbuilder "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/rtctokenbuilder2"
	"github.com/gofiber/fiber/v2"
)

func GenerateToken(c *fiber.Ctx) error {

	req := new(TokenRequest)
	if err := c.BodyParser(req); err != nil {
		return response.Error(c, "Invalid request body", fiber.StatusBadRequest)
	}

	appID := os.Getenv("AGORA_APP_ID")
	appCert := os.Getenv("AGORA_APP_CERTIFICATE")

	if appID == "" || appCert == "" {
		return response.Error(c, "Agora environment variables missing", fiber.StatusInternalServerError)
	}

	token, err := rtctokenbuilder.BuildTokenWithUserAccount(
		appID,
		appCert,
		req.Channel,
		req.Uid,
		rtctokenbuilder.RolePublisher,
		3600,
		3600,
	)

	if err != nil {
		return response.Error(c, "Token generation failed", fiber.StatusInternalServerError)
	}

	return response.Success(c, TokenResponse{Token: token}, "Token generated", fiber.StatusOK)
}
