package cmd

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stakater/Scaler/internal/pkg/cmd/common"
)

func TestRun(t *testing.T) {
	type args struct {
		scalerOptions *common.ScalerOptions
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "Invalid Args",
			args: args{
				scalerOptions: &common.ScalerOptions{
					Max: 1,
					Min: 2,
				},
			},
			wantErr: true,
			err:     fmt.Errorf("Max cannot be less than min"),
		},
		{
			name: "Invalid Provider",
			args: args{
				scalerOptions: &common.ScalerOptions{
					Max:      3,
					Min:      3,
					Desired:  3,
					RoleArn:  "arn:aws:iam::54564",
					Region:   "us-west-2",
					Provider: "invalid",
				},
			},
			wantErr: true,
			err:     fmt.Errorf("Invalid provider specified : invalid"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Run(tt.args.scalerOptions)
			if tt.wantErr {
				if !reflect.DeepEqual(err, tt.err) {
					t.Errorf("Got error = %v, Want Err %v", err, tt.err)
				}
			} else if err != nil {
				t.Errorf("Run(). Not expecting error but Got error = %v", err)
			}
		})
	}
}
