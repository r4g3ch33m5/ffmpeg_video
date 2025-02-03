package cronjob

import (
	"context"
	"reflect"

	"github.com/robfig/cron/v3"
)

// StartCronJob initializes and starts the cron job for creating daily folders
type Task interface {
	GetSchedule() string
	PreProcess() (any, error)
	Process() (any, error)
	PostProcess() (any, error)
}

type BaseTask struct {
}

func (*BaseTask) GetSchedule() string {
	return "0 0 */1 * *"
}

func (*BaseTask) PreProcess() (any, error) {
	return nil, nil
}
func (*BaseTask) PostProcess() (any, error) {
	return nil, nil
}

type CronManager struct {
	tasks []Task
	cron  []*cron.Cron
}

func NewCronManager() *CronManager {
	return &CronManager{
		tasks: make([]Task, 0),
	}
}

func (c *CronManager) RegisterTask(t Task) {
	c.tasks = append(c.tasks, t)
	valChan, errChan := make(chan any), make(chan error)
	for _, t := range c.tasks {
		cronTask := cron.New()
		cronTask.AddJob(t.GetSchedule(), TaskToCronjob(t, valChan, errChan))
	}
}

type job struct {
	Task
	valChan chan any
	errChan chan error
}

func (j *job) Run() {
	val, err := j.Task.PreProcess()
	if err != nil {
		j.errChan <- err
		return
	}
	if reflect.ValueOf(val).IsValid() {
		j.valChan <- val
	}
	val, err = j.Task.Process()
	if err != nil {
		j.errChan <- err
		return
	}
	if reflect.ValueOf(val).IsValid() {
		j.valChan <- val
	}
	val, err = j.Task.PostProcess()
	if err != nil {
		j.errChan <- err
		return
	}
	if reflect.ValueOf(val).IsValid() {
		j.valChan <- val
	}
}

func TaskToCronjob(t Task, valChan chan any, errChan chan error) cron.Job {
	return &job{
		Task:    t,
		valChan: valChan,
		errChan: errChan,
	}
}

func (c *CronManager) Start(ctx context.Context) {
	for _, c := range c.cron {
		c.Start()
	}
	for range ctx.Done() {
		for _, c := range c.cron {
			c.Stop()
		}
	}
}
