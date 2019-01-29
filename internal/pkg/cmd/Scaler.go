package cmd

import (
	"io"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/stakater/Scaler/internal/pkg/cmd/common"
	"github.com/stakater/Scaler/internal/pkg/tasks"
)

//NewScalerCommand to run Scaler
func NewScalerCommand(in io.Reader, out, errOut io.Writer) *cobra.Command {
	scalerOptions := &common.ScalerOptions{
		Options: common.Options{
			Out: out,
			Err: errOut,
		},
	}

	cmd := &cobra.Command{
		Use:   "Scaler",
		Short: "A tool which updates Auto scaling groups",
		RunE: func(cmd *cobra.Command, args []string) error {
			scalerOptions.Cmd = cmd
			scalerOptions.Args = args
			return Run(scalerOptions)
		},
	}

	scalerOptions.AddFlags(cmd)

	return cmd
}

//Run - Executes the action of command
func Run(scalerOptions *common.ScalerOptions) error {
	err := scalerOptions.ValidateArgs()
	if err != nil {
		return err
	}

	task, err := tasks.NewTask(scalerOptions)
	if err != nil {
		return err
	}

	err = task.Run()

	if err != nil {
		logrus.Fatal(err)
	}

	return nil
}
