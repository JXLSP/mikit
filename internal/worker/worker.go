package worker

import (
	"context"
	"fmt"
	"mikit/pkg/worker"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	server "github.com/RichardKnop/machinery/v2"
	redisbackend "github.com/RichardKnop/machinery/v2/backends/redis"
	redisbroker "github.com/RichardKnop/machinery/v2/brokers/redis"
	eagerlock "github.com/RichardKnop/machinery/v2/locks/eager"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// 定义marchinery worker的命令行工具
func NewWorkerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "worker",
		RunE: func(cmd *cobra.Command, args []string) error {
			return start()
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}

	cobra.OnInitialize(initMachineryConfig)

	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.cobra.yaml)")

	return cmd
}

// 初始化marchinery的配置
func initMachineryConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(filepath.Join(home, ".mikit"))
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName("machinery")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// 启动marchinery worker
/*
	@param ctx context.Context
	@return *server.Server
	@return error
*/
func startMachineryWorker(ctx context.Context) (*server.Server, error) {
	opts := &worker.MachineryConfig{
		Broker:       viper.GetString("app.broker"),
		Backend:      viper.GetString("app.backend"),
		DefaultQueue: viper.GetString("app.default_queue"),
	}

	cfg := worker.NewMachineryConfig(opts)

	broker := redisbroker.New(cfg, viper.GetString("app.broker"), "", "", 0)
	backend := redisbackend.New(cfg, viper.GetString("app.backend"), "", "", 1)
	lock := eagerlock.New()
	instance := server.NewServer(cfg, broker, backend, lock)

	wk := instance.NewWorker(viper.GetString("app.default_queue"), viper.GetInt("app.worker_num"))

	go func() {
		// 使用context控制worker生命周期
		go func() {
			<-ctx.Done()
			fmt.Println("Context已取消，worker正在停止...")
		}()
		wk.Launch()
	}()

	return instance, nil
}

// 启动worker
func start() error {
	// 创建上下文
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 启动worker
	instance, err := startMachineryWorker(ctx)
	if err != nil {
		return err
	}

	// 设置信号处理
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// 等待信号
	fmt.Println("Worker启动成功，按Ctrl+C退出")
	s := <-sigChan
	fmt.Printf("收到信号 %s，正在优雅关闭...\n", s)

	// 取消上下文，通知所有worker停止
	cancel()

	// 使用instance停止broker消费任务
	fmt.Println("正在停止broker消费任务...")
	instance.GetBroker().StopConsuming()

	// 等待一段时间，让worker有机会完成当前任务
	fmt.Println("等待worker完成当前任务...")
	time.Sleep(3 * time.Second)

	fmt.Println("Worker已安全退出")
	return nil
}
