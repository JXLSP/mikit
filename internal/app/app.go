package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"mikit/internal/app/controller"
	"mikit/internal/app/store"
	"mikit/pkg/db"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var CFG string

func NewAppCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "mikit",
		Short:        "在线数据库弱口令和未授权检测程序",
		Long:         "适用于懒人的，减少命令行的输入，提升效率的，在线数据库弱口令和未授权检测程序",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return start()
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, item := range args {
				if len(item) > 0 {
					return fmt.Errorf("")
				}
			}
			return nil
		},
	}
	cobra.OnInitialize(initConfig)

	return cmd
}

func initConfig() {
	// 从配置文件读取数据库配置
	viper.SetConfigFile(CFG)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}
}

func start() error {
	// 初始化数据库连接
	dbInstance, err := initDatabase()
	if err != nil {
		return fmt.Errorf("初始化数据库失败: %v", err)
	}

	// 创建Store实例
	storeInstance := store.NewStore(dbInstance)

	// 初始化Gin引擎
	g := gin.New()
	gin.SetMode(viper.GetString("app.mode"))

	// 注册路由，并将Store实例注入到控制器
	registerRoutes(g, storeInstance)

	// 启动HTTP服务器
	srv := startHttpServer(g)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	return nil
}

func startHttpServer(g *gin.Engine) *http.Server {
	srv := &http.Server{
		Addr:    viper.GetString("app.port"),
		Handler: g,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	return srv
}

// 初始化数据库连接
func initDatabase() (*gorm.DB, error) {
	opts := &db.PostgresOptions{
		DBName: viper.GetString("database.name"),
		DBPass: viper.GetString("database.password"),
		DBUser: viper.GetString("database.username"),
		DBHost: viper.GetString("database.host"),
	}

	return db.NewPostgresConnection(opts)
}

// 注册路由并注入依赖
func registerRoutes(g *gin.Engine, storeInstance *store.Store) {
	// 创建控制器实例，注入Store依赖
	tasksController := controller.NewTasksController(storeInstance)

	// API路由组
	api := g.Group("/api")
	{
		// 任务相关路由
		tasks := api.Group("/tasks")
		{
			tasks.POST("/", tasksController.CreateTask)
			tasks.GET("/:id", tasksController.GetTask)
			tasks.DELETE("/:id", tasksController.DeleteTask)
		}
	}
}
