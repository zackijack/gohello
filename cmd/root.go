package cmd

import (
	"fmt"
	"os"

	"github.com/kitabisa/perkakas/v2/log"
	"github.com/spf13/cobra"
	"github.com/zackijack/gohello/config"
	"github.com/zackijack/gohello/internal/app/commons"
	"github.com/zackijack/gohello/internal/app/repository"
	"github.com/zackijack/gohello/internal/app/server"
	"github.com/zackijack/gohello/internal/app/service"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gohello",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
			examples and usage of using your application.`,
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
}

func start() {
	cfg := config.Config()
	logger := log.NewLogger("gohello")

	opt := commons.Options{
		Config: cfg,
		Logger: logger,
	}

	repo := wiringRepository(repository.Option{
		Options: opt,
	})

	service := wiringService(service.Option{
		Options:    opt,
		Repository: repo,
	})

	server := server.NewServer(opt, service)

	// run app
	server.StartApp()
}

func wiringRepository(repoOption repository.Option) *repository.Repository {
	// wiring up all your repos here

	repo := repository.Repository{}

	return &repo
}

func wiringService(serviceOption service.Option) *service.Services {
	// wiring up all services
	hc := service.NewHealthCheck(serviceOption)
	hl := service.NewHelloService(serviceOption)

	svc := service.Services{
		HealthCheck: hc,
		Hello:       hl,
	}

	return &svc
}
