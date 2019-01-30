package providers

import (
	"reflect"
	"testing"

	"github.com/stakater/Scaler/internal/pkg/providers/aws"
)

func TestMapToProvider(t *testing.T) {
	type args struct {
		providerName string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Can get valid provider",
			args: args{
				providerName: "aws",
			},
			want: &aws.Aws{},
		},
		{
			name: "Cannot get invalid provider",
			args: args{
				providerName: "invalid",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapToProvider(tt.args.providerName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapToProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}
