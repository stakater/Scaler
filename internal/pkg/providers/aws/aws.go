package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/sirupsen/logrus"
	"github.com/stakater/Scaler/internal/pkg/cmd/common"
)

// Aws provider class implementing the Provider interface
type Aws struct {
	scalerName      *string
	roleArn         string
	region          *string
	maxSize         *int64
	minSize         *int64
	desiredCapacity *int64
}

//GetName - Returns name of Aws Provider
func (a *Aws) GetName() string {
	return "Amazon Web Services"
}

// Init initializes the Aws Provider Configuration like Access Token and Reion
func (a *Aws) Init(scalerOptions *common.ScalerOptions) error {
	a.scalerName = &scalerOptions.ScalerName
	a.roleArn = scalerOptions.RoleArn
	a.region = &scalerOptions.Region
	a.maxSize = &scalerOptions.Max
	a.minSize = &scalerOptions.Min
	a.desiredCapacity = &scalerOptions.Desired
	return nil
}

// Scale - Modifies auto scaling group.
func (a *Aws) Scale() error {

	// Initial credentials loaded from SDK's default credential chain. Such as
	// the environment, shared credentials (~/.aws/credentials), or EC2 Instance
	// Role. These credentials will be used to to make the STS Assume Role API.
	session, err := session.NewSession()
	if err != nil {
		return err
	}

	// Create the credentials from AssumeRoleProvider to assume the role
	roleCredentials := stscreds.NewCredentials(session, a.roleArn)

	err = a.updateAutoScalingGroup(session, roleCredentials)

	return err
}

func (a *Aws) updateAutoScalingGroup(session *session.Session, credentials *credentials.Credentials) error {

	// Create an Auto Scaling service client.
	asgClient := autoscaling.New(session, &aws.Config{
		Credentials: credentials,
		Region:      a.region,
	})

	_, err := asgClient.UpdateAutoScalingGroup(&autoscaling.UpdateAutoScalingGroupInput{
		AutoScalingGroupName: a.scalerName,
		MaxSize:              a.maxSize,
		MinSize:              a.minSize,
		DesiredCapacity:      a.desiredCapacity,
	})

	if err != nil {
		return err
	}

	logrus.Infof("Successfully modified auto scaling group : %s \n max: %d, min: %d, desired: %d",
		*a.scalerName, *a.maxSize, *a.minSize, *a.desiredCapacity)

	return nil
}
