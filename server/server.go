package main

// Serves two purposes:
// 1. Runs the API server
// 2. Starts the service that always looks for new syncs to sync
func main() {
	go StartApi()
	go StartSyncBridge()

	// Server should never die
	select {}
}
