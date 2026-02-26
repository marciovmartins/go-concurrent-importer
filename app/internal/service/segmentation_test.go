package service

import (
	"go-concurrent-importer/internal/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessRecord(t *testing.T) {
	// arrange
	db := testutil.NewDbFake()
	repo := testutil.NewFakeRepository(db)
	srvc := NewSegmentation(repo)

	records := [][]string{
		{
			"1",
			"PATIENT",
			"João Silva",
			`{""age"":45,""gender"":""M"",""risk"":""high""}`,
		},
	}

	expectedCount := db.SegsCounter + 1

	// act
	_, errs := srvc.ProcessBatch(records)

	// assert
	assert.Empty(t, errs, "não deveria haver erros")
	assert.Equal(t, expectedCount, db.SegsCounter)
}
