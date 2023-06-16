package job_service

import (
	"github.com/go-co-op/gocron"
	"github.com/lowl11/lazylog/log"
	"github.com/lowl11/owl/data/domain"
)

func (job *Job) Run() {
	for _, jobFunc := range job.runFuncList {
		if _, err := job.scheduler.Every(jobFunc.EverySeconds).Second().Do(jobFunc.ExecutableFunc); err != nil {
			log.Error(err, "Run job error")
		}
	}

	job.scheduler.StartAsync()
	job.started = true
}

func (job *Job) Cron(expression string, executableFunc func()) *Job {
	if expression == "" {
		return job
	}

	job.runFuncList = append(job.runFuncList, domain.JobFunc{
		IsCron:         true,
		CronExpression: expression,
		ExecutableFunc: executableFunc,
	})
	return job
}

func (job *Job) Every(seconds int, executableFunc func()) *Job {
	if seconds <= 0 {
		return job
	}

	job.runFuncList = append(job.runFuncList, domain.JobFunc{
		IsEvery:        true,
		EverySeconds:   seconds,
		ExecutableFunc: executableFunc,
	})
	return job
}

func (job *Job) Scheduler() *gocron.Scheduler {
	if job.started {
		panic("Scheduler called after running")
	}

	return job.scheduler
}
