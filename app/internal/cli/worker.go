package cli

import (
	"sync"
)

func (ch *cliHandler) worker(
	batchSize int,
	recordsChan <-chan  []string,
	errorsChan chan<- []error,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	batch := make([][]string, 0, batchSize)

	for record := range recordsChan {
		batch = append(batch, record)

		if len(batch) >= batchSize {
			// saves batch
			_, errs := ch.cliApp.SegmentationService.ProcessBatch(batch)
			if len(errs) > 0 {
				errorsChan <- errs
			}
			batch = batch[:0]
		}
	}

	if len(batch) > 0 {
		_, errs := ch.cliApp.SegmentationService.ProcessBatch(batch)
		if len(errs) > 0 {
			errorsChan <- errs
		}
	}
}