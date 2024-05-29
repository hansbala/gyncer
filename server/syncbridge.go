package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/hansbala/gyncer/core"
)

// I wanted to keep naming unique. This is basically emulating cron.
// All it does is periodically scan the database for syncs to be performed
// and perform them.
func StartSyncBridge() {
	metaSyncFrequency, err := getMetaSyncFrequency()
	if err != nil {
		panic(err)
	}

	// run runSyncJobs every metaSyncFrequency minutes then sleep for metaSyncFrequency minutes
	for {
		runSyncJobs()
		sleepTime := metaSyncFrequency * 60
		fmt.Printf("sleeping for %d minutes\n", metaSyncFrequency)
		sleepTimeInt64 := int64(sleepTime)
		sleepTimeDuration := time.Duration(sleepTimeInt64) * time.Second
		time.Sleep(sleepTimeDuration)
	}
}

// returns sync frequency in minutes (read from environment variable)
func getMetaSyncFrequency() (int, error) {
	syncFrequency := os.Getenv("GYNCER_META_SYNC_FREQUENCY")
	syncFrequencyInt, err := strconv.Atoi(syncFrequency)
	if err != nil {
		return -1, err
	}
	return syncFrequencyInt, nil
}

func runSyncJobs() {
	fmt.Println(core.DatasourceSpotify)
	fmt.Println("running sync jobs")
}
