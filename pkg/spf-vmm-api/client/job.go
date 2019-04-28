package client

import (
	"errors"
	"github.com/badstreff/terraform-provider-vmm/pkg/spf-vmm-api/getting"
	"github.com/badstreff/terraform-provider-vmm/pkg/spf-vmm-api/updating"
	"time"
)

// JobService service
type JobService struct {
	getting.JobGetter
	updating.JobUpdater
}

// WaitForJobToComplete blocks until the job is completed or the timeout in seconds expires
func (s *JobService) WaitForJobToComplete(id string, timeout int) error {
	job, err := s.GetByID(id)
	if err != nil {
		return err
	}
	timeoutChannel := time.After(time.Duration(timeout) * time.Second)
	tick := time.Tick(1 * time.Second)
	for {
		select {
		case <-timeoutChannel:
			return errors.New("timeout exceeded waiting for job to complete")
		case <-tick:
			time.Sleep(1 * time.Second)
			job, err = s.GetByID(id)
			if err != nil {
				return err
			}
			if job.IsCompleted == nil {
				continue
			}
			if *job.IsCompleted {
				return nil
			}
		}
	}
}
