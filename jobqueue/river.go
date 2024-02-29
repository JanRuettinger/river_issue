package jobqueue

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	worker "river_supavisor_issue/jobqueue/workers"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/riverdriver/riverpgxv5"
)

type JobQueue struct {
	RiverClient *river.Client[pgx.Tx]
	pool        *pgxpool.Pool
	logger      *slog.Logger
}

func New(ctx context.Context, databaseURL string, numberWorkers int, logger *slog.Logger) *JobQueue {

	workers := river.NewWorkers()
	documentProcessorWorker := worker.NewDocumentProcessorWorker(logger)
	if err := river.AddWorkerSafely(workers, documentProcessorWorker); err != nil {
		panic(fmt.Sprintf("Document processor worker not added to river client. Error: %v", err))
	}

	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		panic(fmt.Sprintf("error parsing database URL: %v", err))
	}

	config.MaxConnLifetime = 2 * time.Minute
	config.MaxConnIdleTime = 10 * time.Second
	config.MaxConns = 10
	config.MinConns = 5

	dbPool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		panic(fmt.Sprintf("error connecting to the database: %v", err))
	}

	riverClient, err := river.NewClient(riverpgxv5.New(dbPool), &river.Config{
		Queues: map[string]river.QueueConfig{
			river.QueueDefault: {MaxWorkers: numberWorkers},
		},
		Workers:              workers,
		JobTimeout:           10 * time.Minute,
		RescueStuckJobsAfter: 24 * time.Hour,
	})
	if err != nil {
		panic(fmt.Sprintf("error creating river client: %v", err))
	}

	return &JobQueue{RiverClient: riverClient, pool: dbPool, logger: logger}
}

func (jq *JobQueue) Start(ctx context.Context) {
	// Run the client inline. All executed jobs will inherit from ctx:
	if err := jq.RiverClient.Start(ctx); err != nil {
		panic(fmt.Sprintf("error starting river client: %v", err))
	}
}

func (jq *JobQueue) EnqueueDocumentProcessor(ctx context.Context, number int) (int64, error) {

	jobRow, err := jq.RiverClient.Insert(ctx, &worker.DocumentProcessorArgs{Number: number}, &river.InsertOpts{MaxAttempts: 5})

	if err != nil {
		// handle error
		jq.logger.Error("error inserting job: ", err)
		return 0, err
	}
	return jobRow.ID, nil
}
