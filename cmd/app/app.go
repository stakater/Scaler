package app

import (
	"os"

	"github.com/stakater/Scaler/internal/pkg/cmd"
)

// Run runs the Scaler command
func Run() error {
	cmd := cmd.NewScalerCommand(os.Stdin, os.Stdout, os.Stderr)
	return cmd.Execute()
}
