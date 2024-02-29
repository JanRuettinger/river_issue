package worker

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/riverqueue/river"
)

type DocumentProcessorArgs struct {
	Number int `json:"Number"`
}

func (DocumentProcessorArgs) Kind() string { return "document_processor" }

type DocumentProcessorWorker struct {
	// An embedded WorkerDefaults sets up default methods to fulfill the rest of
	// the Worker interface:
	river.WorkerDefaults[DocumentProcessorArgs]
	logger *slog.Logger
}

func (w *DocumentProcessorWorker) Work(ctx context.Context, job *river.Job[DocumentProcessorArgs]) error {

	number := job.Args.Number

	for i := 0; i < number; i++ {
		w.logger.Info(fmt.Sprintf("Processing document %d", i))
		time.Sleep(1 * time.Second)
	}

	return nil
}

func NewDocumentProcessorWorker(logger *slog.Logger) *DocumentProcessorWorker {
	return &DocumentProcessorWorker{logger: logger}
}
