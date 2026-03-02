package cli

import (
	"sync"
)

func (h *Handler) worker(
	recordsChan <-chan []string,
	errorsChan chan<- []error,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	batch := make([][]string, 0, h.cfg.BatchSize)

	for record := range recordsChan {
		batch = append(batch, record)

		if len(batch) >= h.cfg.BatchSize {
			// saves batch
			_, errs := h.segmentationService.ProcessBatch(batch)
			if len(errs) > 0 {
				errorsChan <- errs
			}
			batch = batch[:0]
		}
	}

	if len(batch) > 0 {
		_, errs := h.segmentationService.ProcessBatch(batch)
		if len(errs) > 0 {
			errorsChan <- errs
		}
	}
}
