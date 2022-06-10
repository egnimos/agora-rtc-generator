package controllers

import (
	"net/http"
	"time"

	"github.com/AgoraIO-Community/go-tokenbuilder/rtctokenbuilder"
	"github.com/egnimos/agora-rtc-generator/src/app_env"
	"github.com/egnimos/agora-rtc-generator/src/domain"
	"github.com/egnimos/agora-rtc-generator/src/services"
	"github.com/gin-gonic/gin"
)

var (
	RTCController RTCControllerInterface = &rtccontroller{}
)

type RTCControllerInterface interface {
	GenerateRTCToken(ctx *gin.Context)
}

type rtccontroller struct{}

func parseRtcParams(c *gin.Context) (channelName, tokentype, uidStr string, role rtctokenbuilder.Role, expireTimestamp uint32, err *domain.RestError) {
	channelName = c.Param("channel")
	roleStr := c.Param("role")
	tokentype = c.Param("token_type")
	uidStr = c.Param("uid")

	if roleStr == "publisher" {
		role = rtctokenbuilder.RolePublisher
	} else {
		role = rtctokenbuilder.RoleSubscriber
	}

	// set timestamps
	expireTimeInSeconds := uint32(3600)
	currentTimestamp := uint32(time.Now().UTC().Unix())
	expireTimestamp = currentTimestamp + expireTimeInSeconds

	return channelName, tokentype, uidStr, role, expireTimestamp, nil
}

func (rtc *rtccontroller) GenerateRTCToken(ctx *gin.Context) {
	//"rtc/:channel/:role/:token_type/:uid/"
	// get param values
	channelName, tokenType, uidStr, role, expireTimestamp, err := parseRtcParams(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	appID, _ := app_env.InitEnv.GetAppId()
	appCertificate, _ := app_env.InitEnv.GetAppCert()
	result, genErr := services.GenToken.GenerateRTCToken(tokenType, appID, appCertificate, channelName, uidStr, role, expireTimestamp)
	if genErr != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}
