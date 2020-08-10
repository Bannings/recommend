package timed_task

import (
	"time"
)

func ReloadTimedTaskData() {
	reloadNewUserRuleData()
	reloadOldUserRuleData()
	ReloadRecommendRuleDataWxapp()
	UpdateTagMap()
	return
}

func ExcuteTimedTask() {
	ticker := time.NewTicker(900 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			ReloadTimedTaskData()
			ReloadBannerData()
		}
	}
}
