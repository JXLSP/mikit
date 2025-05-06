package worker

import "github.com/RichardKnop/machinery/v2/config"

// 用于machinery的配置
type MachineryConfig struct {
	Broker       string
	Backend      string
	DefaultQueue string
}

// 新建一个machinery的配置
func NewMachineryConfig(opts *MachineryConfig) *config.Config {
	return &config.Config{
		Broker:        opts.Broker,
		DefaultQueue:  opts.DefaultQueue,
		ResultBackend: opts.Backend,
	}
}
