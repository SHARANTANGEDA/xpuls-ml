package cmd

import (
	"context"
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/xpuls-com/xpuls-ml/config"
	"github.com/xpuls-com/xpuls-ml/crons"
	"github.com/xpuls-com/xpuls-ml/dto"
	"github.com/xpuls-com/xpuls-ml/routes"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type ServeOption struct {
	ConfigPath string
}

func (opt *ServeOption) Validate(ctx context.Context) error {
	return nil
}

func (opt *ServeOption) Complete(ctx context.Context, args []string, argsLenAtDash int) error {
	return nil
}

func (opt *ServeOption) addCronDaemons(ctx context.Context) {
	s := gocron.NewScheduler(time.UTC).SingletonMode()
	logger := logrus.New().WithField("cron", "sync env")
	// Add cron every 30 seconds for processing llm tracing queue
	_, err := s.Every(30).Second().Do(func() {
		logger.Info("starting processor for llm tracing queue")
		err := crons.AddProcessLLMTracingQueue(ctx)
		if err != nil {
			return
		}
	})
	if err != nil {
		return
	}

	// Add cron every 2 minutes for deleting processed llm traces
	_, err = s.Every(2).Minute().Do(func() {
		logger.Info("starting cleaner for deleting processed queue items")

		err := crons.DeleteProcessedQueueItems(ctx)
		if err != nil {
			return
		}
	})
	if err != nil {
		return
	}
	s.StartAsync()
}

func (opt *ServeOption) Run(ctx context.Context, args []string) error {
	if !GlobalCommandOption.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	content, err := os.ReadFile(opt.ConfigPath)
	if err != nil {
		return errors.Wrapf(err, "read config file: %s", opt.ConfigPath)
	}

	err = yaml.Unmarshal(content, config.ServerConfig)
	if err != nil {
		return errors.Wrapf(err, "unmarshal config file: %s", opt.ConfigPath)
	}

	err = config.CreateServerConfig()
	if err != nil {
		return errors.Wrapf(err, "populate config file: %s", opt.ConfigPath)
	}

	err = dto.MigrateUp()
	if err != nil {
		return errors.Wrap(err, "migrate up db")
	}

	// Add cron daemons
	opt.addCronDaemons(ctx)

	// nolint: contextcheck
	router, err := routes.NewRouter()
	if err != nil {
		return err
	}

	readHeaderTimeout := 10 * time.Second

	logrus.Infof("listening on 0.0.0.0:%d", config.ServerConfig.Port)

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", config.ServerConfig.Port),
		Handler:           router,
		ReadHeaderTimeout: readHeaderTimeout,
	}
	return srv.ListenAndServe()
}

func getServeCmd() *cobra.Command {
	var opt ServeOption
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "run xpuls-ml api server",
		Long:  "",
		RunE:  config.MakeRunE(&opt),
	}
	cmd.Flags().StringVarP(&opt.ConfigPath, "config", "c", "./xpuls-config.yaml", "")
	return cmd
}
