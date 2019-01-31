package common

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/cobra"
)

func TestScalerOptions_AddFlags(t *testing.T) {
	type fields struct {
		Options    Options
		ScalerName string
		Provider   string
		RoleArn    string
		Region     string
		Max        int64
		Min        int64
		Desired    int64
	}
	type args struct {
		cmd *cobra.Command
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Add Flags works",
			fields: fields{},
			args: args{
				cmd: &cobra.Command{
					Use:   "Scaler",
					Short: "A tool which updates Auto scaling groups",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scalerOptions := &ScalerOptions{
				Options:    tt.fields.Options,
				ScalerName: tt.fields.ScalerName,
				Provider:   tt.fields.Provider,
				RoleArn:    tt.fields.RoleArn,
				Region:     tt.fields.Region,
				Max:        tt.fields.Max,
				Min:        tt.fields.Min,
				Desired:    tt.fields.Desired,
			}
			scalerOptions.AddFlags(tt.args.cmd)
		})
	}
}

func TestScalerOptions_ValidateArgs(t *testing.T) {
	type fields struct {
		Options    Options
		ScalerName string
		Provider   string
		RoleArn    string
		Region     string
		Max        int64
		Min        int64
		Desired    int64
	}
	tests := []struct {
		name   string
		fields fields
		want   error
	}{
		{
			name: "Validate Args - max < min",
			fields: fields{
				Max: 2,
				Min: 3,
			},
			want: fmt.Errorf("Max cannot be less than min"),
		},
		{
			name: "Validate Args - max < desired",
			fields: fields{
				Max:     3,
				Min:     3,
				Desired: 4,
			},
			want: fmt.Errorf("Max cannot be less than desired"),
		},
		{
			name: "Validate Args - desired < min",
			fields: fields{
				Max:     3,
				Min:     3,
				Desired: 2,
			},
			want: fmt.Errorf("Desired cannot be less than min"),
		},
		{
			name: "Validate Args - role ARN",
			fields: fields{
				Max:     3,
				Min:     3,
				Desired: 3,
				RoleArn: "invalid-arn:aws:iam::randomID",
			},
			want: fmt.Errorf("Role Arn must start with : 'arn:aws:iam::'"),
		},
		{
			name: "Validate Args - region",
			fields: fields{
				Max:     3,
				Min:     3,
				Desired: 3,
				RoleArn: "arn:aws:iam::randomID",
				Region:  "Ireland",
			},
			want: fmt.Errorf("Not a valid aws region. Valid regions match expression : \"^(us|eu|ap|sa|ca)\\-\\w+\\-\\d+$\""),
		},
		{
			name: "Validate Args - valid args",
			fields: fields{
				Max:     3,
				Min:     3,
				Desired: 3,
				RoleArn: "arn:aws:iam::54564",
				Region:  "us-west-2",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scalerOptions := &ScalerOptions{
				Options:    tt.fields.Options,
				ScalerName: tt.fields.ScalerName,
				Provider:   tt.fields.Provider,
				RoleArn:    tt.fields.RoleArn,
				Region:     tt.fields.Region,
				Max:        tt.fields.Max,
				Min:        tt.fields.Min,
				Desired:    tt.fields.Desired,
			}
			got := scalerOptions.ValidateArgs()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScalerOptions.ValidateArgs() error = %v, wantErr %v", got, tt.want)
			}
		})
	}
}
