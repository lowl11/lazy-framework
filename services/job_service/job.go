package job_service

import (
	"github.com/go-co-op/gocron"
	"github.com/lowl11/lazy-framework/data/domain"
	"time"
)

type Job struct {
	scheduler   *gocron.Scheduler
	runFuncList []domain.JobFunc

	started bool
}

func New() *Job {
	return &Job{
		scheduler:   gocron.NewScheduler(time.UTC),
		runFuncList: make([]domain.JobFunc, 0),
	}
}
