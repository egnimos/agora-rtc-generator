package app

import "github.com/egnimos/agora-rtc-generator/src/controllers";

func mapurl() {
		router.GET("/rtc/:channel/publisher/uid/:uid", controllers.RTCController.GenerateRTCToken)
}