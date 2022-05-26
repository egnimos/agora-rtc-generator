package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/egnimos/agora-rtc-generator/src/app_env"
	rtctokenbuilder "github.com/egnimos/agora-rtc-generator/src/services/rtc_token_builder"
	"github.com/gin-gonic/gin"
)

var (
	RTCController RTCControllerInterface = &rtccontroller{}
)

type RTCControllerInterface interface {
	GenerateRTCToken(ctx *gin.Context)
}

type rtccontroller struct{}

func (rtc *rtccontroller) GenerateRTCToken(ctx *gin.Context) {
	//get the channel name
	channel := ctx.Param("channel")
	//get the uid
	u, err := strconv.ParseInt(ctx.Param("uid"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	uid := uint(u)

	//generate the token
	appID, _ := app_env.InitEnv.GetAppId()
	appCertificate, _ := app_env.InitEnv.GetAppCert()
	expireTimeInSeconds := uint32(3600)
	currentTimestamp := uint32(time.Now().UTC().Unix())
	expireTimestamp := currentTimestamp + expireTimeInSeconds

	result, err := rtctokenbuilder.BuildTokenWithUID(appID, appCertificate, channel, uint32(uid), rtctokenbuilder.RoleAttendee, expireTimestamp)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}
