package holiday

import (
	"testing"
	"time"
)

func TestLoadData(t *testing.T) {
	index := mustLoadBuiltInData()

	if len(index) == 0 {
		t.Fatalf("IndexByDate must not be empty")
	}

	for idx, holiday := range index {
		if len(holiday.States) == 0 {
			t.Fatalf("error at %v: holidays without states should be checked and removed", idx)
		}
		if len(holiday.Name) == 0 {
			t.Fatalf("error at %v: holidays without a name should be checked and removed", idx)
		}
		if _, err := time.Parse("2006-01-02", idx); err != nil {
			t.Fatalf("error at %v: holidays with invalid dates should be checked and removed", idx)
		}
	}
}
