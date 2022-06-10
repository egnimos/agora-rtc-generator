package services

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/AgoraIO-Community/go-tokenbuilder/rtctokenbuilder"
	"github.com/egnimos/agora-rtc-generator/src/domain"
)

var (
	GenToken GenerateTokenInterface = &genToken{}
)

type GenerateTokenInterface interface {
	GenerateRTCToken(tokenType string, appId string, appCert string, channelName string, uidStr string, role rtctokenbuilder.Role, expireTimestamp uint32) (*domain.RestError, map[string]interface{}) 
}

type genToken struct{}

func (gt *genToken) GenerateRTCToken(tokenType string, appId string, appCert string, channelName string, uidStr string, role rtctokenbuilder.Role, expireTimestamp uint32) (*domain.RestError, map[string]interface{}) {
	if tokenType == "userAccount" {
		log.Printf("Building Token with userAccount: %s\n", uidStr)
		rtcToken, err := rtctokenbuilder.BuildTokenWithUserAccount(appId, appCert, channelName, uidStr, role, expireTimestamp)
		if err != nil {
			return &domain.RestError{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
			}, nil
		}
		return nil, map[string]interface{}{
			"rtcToken": rtcToken,
		}
	} else if tokenType == "uid" {
		uid64, parseErr := strconv.ParseUint(uidStr, 10, 64)
		// check if conversion fails
		if parseErr != nil {
			err := fmt.Sprintf("failed to parse uidStr: %s, to uint causing error: %s", uidStr, parseErr)
			return &domain.RestError{
				Message: err,
				Status:  http.StatusBadRequest,
			}, nil
		}

		uid := uint32(uid64) // convert uid from uint64 to uint 32
		log.Printf("Building Token with uid: %d\n", uid)
		rtcToken, err := rtctokenbuilder.BuildTokenWithUID(appId, appCert, channelName, uid, role, expireTimestamp)
		if err != nil {
			return &domain.RestError{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
			}, nil
		}
		return nil, map[string]interface{}{
			"rtcToken": rtcToken,
		}
	} else {
		err := fmt.Sprintf("failed to generate RTC token for Unknown Tokentype: %s", tokenType)
		return &domain.RestError{
			Message: err,
			Status:  http.StatusBadRequest,
		}, nil
	}
}
