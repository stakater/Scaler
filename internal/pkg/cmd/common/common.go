package common

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

// Options contains common options and helper methods
type Options struct {
	Out     io.Writer
	Err     io.Writer
	Cmd     *cobra.Command
	Args    []string
	Verbose bool
}

//ScalerOptions saves values of flags
type ScalerOptions struct {
	Options
	ScalerName string
	Provider   string
	RoleArn    string
	Region     string
	Max        int64
	Min        int64
	Desired    int64
}

//AddFlags - adds required flags to command
func (scalerOptions *ScalerOptions) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&scalerOptions.ScalerName, "scalerName", "n", "", "Name of Auto Scaling group (Required)")
	cmd.Flags().StringVarP(&scalerOptions.Provider, "provider", "p", "", "Cloud provider to user. Valid Values \n 1- aws")
	cmd.Flags().StringVarP(&scalerOptions.RoleArn, "roleArn", "a", "", "Arn of role to assume (Required)")
	cmd.Flags().StringVarP(&scalerOptions.Region, "region", "r", "", "Region in which auto scaling group exists (Required)")
	cmd.Flags().Int64VarP(&scalerOptions.Max, "max", "m", 0, "Maximum no of instances (Required)")
	cmd.Flags().Int64VarP(&scalerOptions.Min, "min", "i", 0, "Minimum no of instances (Required)")
	cmd.Flags().Int64VarP(&scalerOptions.Desired, "desired", "d", 0, "Desired no of instances (Required)")

	cmd.MarkFlagRequired("scalerName")
	cmd.MarkFlagRequired("provider")
	cmd.MarkFlagRequired("roleArn")
	cmd.MarkFlagRequired("region")
	cmd.MarkFlagRequired("max")
	cmd.MarkFlagRequired("min")
	cmd.MarkFlagRequired("desired")

}

//ValidateArgs - Validates the arguments
func (scalerOptions *ScalerOptions) ValidateArgs() error {
	if scalerOptions.Max < scalerOptions.Min {
		return fmt.Errorf("Max cannot be less than min")
	}

	if scalerOptions.Max < scalerOptions.Desired {
		return fmt.Errorf("Max cannot be less than desired")
	}

	if scalerOptions.Desired < scalerOptions.Min {
		return fmt.Errorf("Desired cannot be less than min")
	}

	if !strings.HasPrefix(scalerOptions.RoleArn, "arn:aws:iam::") {
		return fmt.Errorf("Role Arn must start with : 'arn:aws:iam::'")
	}

	reg, _ := regexp.Compile("^(us|eu|ap|sa|ca)\\-\\w+\\-\\d+$")
	if !reg.MatchString(scalerOptions.Region) {
		return fmt.Errorf("Not a valid aws region. Valid regions match expression : \"^(us|eu|ap|sa|ca)\\-\\w+\\-\\d+$\"")
	}

	return nil
}
