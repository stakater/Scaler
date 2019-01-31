package aws

import (
	"testing"

	"github.com/stakater/Scaler/internal/pkg/cmd/common"
)

func TestAws_GetName(t *testing.T) {
	t.Run("Get Provider Name", func(t *testing.T) {
		a := &Aws{}
		if got := a.GetName(); got != "Amazon Web Services" {
			t.Errorf("Aws.GetName() = %v, want Amazon Web Services", got)
		}
	})
}

func TestAws_Init(t *testing.T) {
	t.Run("Can initialize Provider", func(t *testing.T) {
		scalerOptions := &common.ScalerOptions{
			ScalerName: "Test Auto Scaling Name",
			RoleArn:    "arn:aws:iam::54564",
			Region:     "us-west-2",
			Max:        3,
			Min:        3,
			Desired:    3,
		}
		a := &Aws{}
		a.Init(scalerOptions)
		if scalerOptions.ScalerName != *a.scalerName {
			t.Errorf("Expecting Scalar Name: %s, Got: %s", scalerOptions.ScalerName, *a.scalerName)
		}

		if scalerOptions.RoleArn != a.roleArn {
			t.Errorf("Expecting RoleArn: %s, Got: %s", scalerOptions.RoleArn, a.roleArn)
		}

		if scalerOptions.Region != *a.region {
			t.Errorf("Expecting region: %s, Got: %s", scalerOptions.Region, *a.region)
		}

		if scalerOptions.Max != *a.maxSize {
			t.Errorf("Expecting Max: %d, Got: %d", scalerOptions.Max, *a.maxSize)
		}

		if scalerOptions.Min != *a.minSize {
			t.Errorf("Expecting Min: %d, Got: %d", scalerOptions.Min, *a.minSize)
		}

		if scalerOptions.Desired != *a.desiredCapacity {
			t.Errorf("Expecting Min: %d, Got: %d", scalerOptions.Desired, *a.desiredCapacity)
		}
	})
}
