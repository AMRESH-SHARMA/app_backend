package rtc

import (
	"app_backend/internal/response"

	rtctokenbuilder "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/rtctokenbuilder2"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func GenerateToken(c *fiber.Ctx) error {

	req := new(TokenRequest)
	if err := c.BodyParser(req); err != nil {
		return response.Error(c, "Invalid request body", fiber.StatusBadRequest)
	}

	appID := viper.GetString("AGORA.APP_ID")
	appCert := viper.GetString("AGORA.APP_CERTIFICATE")
	tokenExpiry := uint32(viper.GetInt("AGORA.TOKEN_EXPIRY"))

	if appID == "" || appCert == "" || tokenExpiry == 0 {
		return response.Error(c, "Agora environment variables missing", fiber.StatusInternalServerError)
	}

	token, err := rtctokenbuilder.BuildTokenWithUserAccount(
		appID,
		appCert,
		req.Channel, // This is callId
		req.Uid,
		rtctokenbuilder.RolePublisher,
		tokenExpiry,
		tokenExpiry,
	)

	if err != nil {
		return response.Error(c, "Token generation failed", fiber.StatusInternalServerError)
	}

	return response.Success(c, TokenResponse{Token: token}, "Token generated", fiber.StatusOK)
}
