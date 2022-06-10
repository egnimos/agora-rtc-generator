package app

import "github.com/egnimos/agora-rtc-generator/src/controllers"

func mapurl() {
	//ping
	router.GET("/ping", controllers.Ping)
	//generate token
	router.GET("rtc/:channel/:role/:token_type/:uid/", controllers.RTCController.GenerateRTCToken)
}
