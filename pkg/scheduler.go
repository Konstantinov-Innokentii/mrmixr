package pkg

import (
	"time"

	"github.com/go-co-op/gocron"
)

func task() {
	// Get all checks, filter by execution time
	// Start fetch and notify process in separate goroutine
}

func startScheduler() error {
	// defines a new scheduler that schedules and runs jobs
	scheduler := gocron.NewScheduler(time.UTC)
	_, err := scheduler.Cron("*/30 * * * *").Do(task)
	if err != nil {
		return err
	}
	return nil
}
