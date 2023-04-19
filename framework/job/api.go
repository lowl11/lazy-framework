package job

import (
	"github.com/lowl11/lazy-framework/data/domain"
	"github.com/lowl11/lazy-framework/log"
)

func (job *Job) Run() {
	for _, jobFunc := range job.runFuncList {
		if _, err := job.scheduler.Every(jobFunc.EverySeconds).Second().Do(jobFunc.ExecutableFunc); err != nil {
			log.Error(err, "Run job error")
		}
	}

	job.scheduler.StartAsync()
}

func (job *Job) Cron(expression string, executableFunc func()) *Job {
	job.runFuncList = append(job.runFuncList, domain.JobFunc{
		IsCron:         true,
		CronExpression: expression,
		ExecutableFunc: executableFunc,
	})
	return job
}

func (job *Job) Every(seconds int, executableFunc func()) *Job {
	job.runFuncList = append(job.runFuncList, domain.JobFunc{
		IsEvery:        true,
		EverySeconds:   seconds,
		ExecutableFunc: executableFunc,
	})
	return job
}
