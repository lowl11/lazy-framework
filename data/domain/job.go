package domain

type JobFunc struct {
	IsEvery bool
	IsCron  bool

	CronExpression string
	EverySeconds   int

	ExecutableFunc func()
}
