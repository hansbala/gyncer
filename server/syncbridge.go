package main

import (
	"fmt"
	"time"

	"github.com/hansbala/gyncer/config"
	"github.com/hansbala/gyncer/core"
)

// I wanted to keep naming unique. This is basically emulating cron.
// All it does is periodically scan the database for syncs to be performed
// and perform them.
func StartSyncBridge() {
	config := config.GetConfig()
	syncFrequencyMinutes := config.Server.MetaSyncFrequency
	// fmt.Println(syncFrequencyMinutes)

	// run sync jobs and sleep for syncFrequencyMinutes
	for {
		runSyncJobs()
		sleepTime := time.Duration(syncFrequencyMinutes) * time.Minute
		time.Sleep(sleepTime)
	}
}

func runSyncJobs() {
	fmt.Println(core.DatasourceSpotify)
	fmt.Println("running sync jobs")
}
