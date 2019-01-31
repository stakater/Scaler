package providers

import (
	"github.com/stakater/Scaler/internal/pkg/cmd/common"
	"github.com/stakater/Scaler/internal/pkg/providers/aws"
)

// Provider interface so that providers like aws, google cloud can implement this
type Provider interface {
	Init(options *common.ScalerOptions) error
	Scale() error
}

// MapToProvider maps the IP provider name to the actual IpProvider type
func MapToProvider(providerName string) Provider {
	ipProvider, ok := providerMap[providerName]
	if !ok {
		return nil
	}
	return ipProvider
}

var providerMap = map[string]Provider{
	"aws": &aws.Aws{},
}
