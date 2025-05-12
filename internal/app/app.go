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

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

func initConfig() {}

func start() error {
	g := gin.New()
	gin.SetMode("app.mode")
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
