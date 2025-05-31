package main

import (
	"time"
	// db "bdoItemFetcher/db"
	// getitems "bdoItemFetcher/getitems"
	updater "bdoItemFetcher/updater"
)

func main() {
	for {
		updater.Updater()
		time.Sleep(6 * time.Hour)
	}
}