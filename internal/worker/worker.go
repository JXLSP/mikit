package worker

import (
	"github.com/spf13/cobra"
)

func NewWorkerCommand() *cobra.Command {
	return &cobra.Command{
		Use: "worker",
		RunE: func(cmd *cobra.Command, args []string) error {
			return start()
		},
	}
}

func start() error {
	return nil
}
