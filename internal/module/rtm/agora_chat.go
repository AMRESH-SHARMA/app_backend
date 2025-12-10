package rtm

import (
	"fmt"

	"github.com/spf13/viper"
)

func GenerateRTMToken(userID string) (string, error) {
	appID := viper.GetString("AGORA.CHAT_APP_ID")
	appCert := viper.GetString("AGORA.CHAT_APP_CERT")

	if appID == "" || appCert == "" {
		return "", fmt.Errorf("Agora Chat credentials missing")
	}

	// pseudo generate token; implement using Agora Chat REST API
	token := fmt.Sprintf("token_%s_%s", appID, userID)

	return token, nil
}
