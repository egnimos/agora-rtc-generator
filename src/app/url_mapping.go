package app

import "github.com/egnimos/agora-rtc-generator/src/controllers";

func mapurl() {
	router.POST("/rtc/:channel/publisher/uid/:uid", controllers.RTCController.GenerateRTCToken)
}