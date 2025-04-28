package app

import "github.com/spf13/cobra"

func NewAppCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "mikit",
		Short:        "在线数据库弱口令和未授权检测程序",
		Long:         "适用于懒人的，减少命令行的输入，提升效率的，在线数据库弱口令和未授权检测程序",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return start()
		},
	}

	return cmd
}

func start() error {
	return nil
}
