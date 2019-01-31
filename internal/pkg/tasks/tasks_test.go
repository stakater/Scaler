package tasks

import (
	"fmt"
	"testing"

	"github.com/stakater/Scaler/internal/pkg/cmd/common"
)

func TestNewTask(t *testing.T) {
	scalerOptions := &common.ScalerOptions{
		Provider:   "aws",
		ScalerName: "Test Auto Scaling Name",
		RoleArn:    "arn:aws:iam::54564",
		Region:     "us-west-2",
		Max:        3,
		Min:        3,
		Desired:    3,
	}

	type args struct {
		scalerOptions *common.ScalerOptions
	}
	tests := []struct {
		name         string
		args         args
		wantProvider bool
		wantErr      bool
		err          error
	}{
		{
			name: "Can create task with valid provider",
			args: args{
				scalerOptions: scalerOptions,
			},
			wantProvider: true,
			wantErr:      false,
		},
		{
			name: "Cannot create task with invalid provider",
			args: args{
				scalerOptions: &common.ScalerOptions{
					Provider: "Invalid Provider",
				},
			},
			wantProvider: false,
			wantErr:      true,
			err:          fmt.Errorf("Invalid provider specified : Invalid Provider"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTask(tt.args.scalerOptions)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got != nil && got.provider != nil) != tt.wantProvider {
				t.Errorf("NewTask() Got Provider: %v, Wanted Provider: %v",
					(got != nil && got.provider != nil), tt.wantProvider)
			}
		})
	}
}
