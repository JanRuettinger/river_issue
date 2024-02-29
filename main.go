package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"river_supavisor_issue/jobqueue"

	"github.com/golang-cz/devslog"
)

const NUMBER_OF_WORKERS = 10

func main() {

	ctx := context.Background()
	w := os.Stdout

	// slog.HandlerOptions
	slogOpts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}

	// new logger with options
	opts := &devslog.Options{
		HandlerOptions:    slogOpts,
		MaxSlicePrintSize: 4,
		SortKeys:          true,
	}

	logger := slog.New(devslog.NewHandler(w, opts))

	databaseURL := "postgres://postgres.XXXXX@aws-0-us-west-1.pooler.supabase.com:5432/postgres"

	jobQueue := jobqueue.New(ctx, databaseURL, NUMBER_OF_WORKERS, logger)
	jobQueue.Start(ctx)
	logger.Info("JobQueue started.")

	time.Sleep(5 * time.Minute)
}
