package holiday

import (
	"fmt"
	"io/ioutil"
	"time"
)

// New returns a new Query instance with built-in holiday data.
func New() *Query {
	return &Query{
		Index: mustLoadBuiltInData(),
	}
}

// Query allows to check for holidays
type Query struct {
	Index IndexByDate
}

// IsHoliday returns true if a holiday was found for the given date
func (q Query) IsHoliday(date time.Time, state string) bool {
	return isHoliday(q.Index, date, state)
}

// AddFileToIndex adds File data from the given filepath to the actual Query index.
// Existing entries in the index are overwritten.
func (q Query) AddFileToIndex(file string) error {
	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("cannot add file '%s' to index: %w", file, err)
	}

	return addToIndex(fileBytes, q.Index)
}

// IsHoliday returns true if a holiday was found for the given date
func IsHoliday(date time.Time, state string) bool {
	return isHoliday(mustLoadBuiltInData(), date, state)
}

func isHoliday(i IndexByDate, date time.Time, state string) bool {
	var idx = date.Format(dateLayout)

	if entry, ok := i[idx]; ok {
		if len(state) == 0 {
			return true
		}

		if entry.States.match(state) {
			return true
		}
	}

	return false
}

func (s States) match(state string) bool {
	for i := range s {
		if s[i] == state {
			return true
		}
	}

	return false
}
