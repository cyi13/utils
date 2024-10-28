package crontab

import (
	"github.com/robfig/cron/v3"
)

var defaultCron *Service

type Service struct {
	cron *cron.Cron
}

func init() {
	defaultCron = &Service{
		cron: cron.New(cron.WithSeconds()),
	}
	defaultCron.cron.Start()
}

func AddCron(rule string, callback func()) {
	defaultCron.AddCron(rule, callback)
}

func (s *Service) AddCron(rule string, callback func()) {
	s.cron.AddFunc(rule, callback)
}

func (s *Service) Stop() {
	s.cron.Stop()
}
