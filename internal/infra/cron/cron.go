package cron

import (
	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	cron *cron.Cron
}

func New(cron *cron.Cron) *Scheduler {
	return &Scheduler{
		cron: cron,
	}
}

func (s *Scheduler) Start() {
	s.cron.Start()
}

func (s *Scheduler) AddTask(spec string, cmd func()) (cron.EntryID, error) {
	return s.cron.AddFunc(spec, cmd)
}

func (s *Scheduler) Stop() {
	s.cron.Stop()
}
