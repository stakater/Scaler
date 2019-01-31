package tasks

import (
	"fmt"

	"github.com/stakater/Scaler/internal/pkg/cmd/common"
	"github.com/stakater/Scaler/internal/pkg/providers"
)

// Task Modify Auto Scaling Group
type Task struct {
	provider providers.Provider
}

// NewTask for initializing the Task
func NewTask(scalerOptions *common.ScalerOptions) (*Task, error) {
	task := &Task{}

	task.provider = providers.MapToProvider(scalerOptions.Provider)
	if task.provider == nil {
		return nil, fmt.Errorf("Invalid provider specified : %s", scalerOptions.Provider)
	}
	task.provider.Init(scalerOptions)
	return task, nil
}

//Run function for task which handles the logic
func (t *Task) Run() error {
	return t.provider.Scale()
}
