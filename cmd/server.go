package cmd

import (
	"context"
	"github.com/jihanlugas/inventory/config"
	"github.com/jihanlugas/inventory/db"
	"github.com/jihanlugas/inventory/log"
	"github.com/jihanlugas/inventory/router"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run server",
	Long:  "Run server",
	Run: func(cmd *cobra.Command, args []string) {
		runServer()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func runServer() {
	var err error
	log.Run()
	dbpool := db.Initialize()
	systemCleaning := make(chan struct{}, 1)
	var wg sync.WaitGroup
	wg.Add(1)

	go scheduler(systemCleaning, &wg)

	var shutdownCallback = func() {
		defer wg.Done()
		close(systemCleaning)

		log.System.Info().Msg("Cleaning resources!")
		dbpool.Close()
	}

	r := router.Init()

	wg.Add(1)
	go func() {
		r.Server.RegisterOnShutdown(shutdownCallback)
		if config.Environment == config.Production {
			if err = r.Start(":" + config.ListenTo.Port); err != nil && err != http.ErrServerClosed {
				r.Logger.Fatal("Shutting down the server")
			}
		} else {
			if err = r.StartTLS(":"+config.ListenTo.Port, config.CertificateFilePath, config.CertificateKeyFilePath); err != nil && err != http.ErrServerClosed {
				r.Logger.Fatal("Shutting down the server")
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := r.Shutdown(ctx); err != nil {
		r.Logger.Fatal(err)
	}

	wg.Wait()
	log.System.Info().Msg("Main System Shutdown!")
	log.CloseAll()
}

func scheduler(systemCleaning chan struct{}, wg *sync.WaitGroup) {
	now := time.Now()
	tomorrow := now.Add(24 * time.Hour)
	midnight := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 0, 0, 0, 0, time.Local)
	everyMidnight := time.NewTimer(midnight.Sub(now))
	//everyMinute := time.NewTicker(time.Minute)

DailyLoop:
	for {
		select {
		case <-everyMidnight.C:
			everyMidnight.Reset(24 * time.Hour)
			log.ChangeDay()
			runtime.GC()
		//case <-everyMinute.C:
		//resetStatusPembayaran()
		case <-systemCleaning:
			if !everyMidnight.Stop() {
				<-everyMidnight.C
			}
			break DailyLoop
		}
	}
	log.System.Warn().Msg("Scheduler shutdown")
	wg.Done()
}
