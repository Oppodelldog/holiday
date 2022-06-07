package holiday

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"
)

func TestIsHoliday(t *testing.T) {
	var testCases = []struct {
		date  string
		state string
		want  bool
	}{
		{date: "2024-12-27", want: false},
		{date: "2024-12-26", want: true},
		{date: "2024-11-20", want: true},
		{date: "2024-11-20", state: Bavaria, want: false},
		{date: "2024-11-20", state: Saxony, want: true},
	}

	for i, testCase := range testCases {
		date := parseDate(t, testCase.date)
		var got = IsHoliday(date, testCase.state)

		if testCase.want != got {
			t.Fatalf("test case (%v): want: %v, got: %v", i, testCase.want, got)
		}
	}
}

func TestQuery_AddFileToIndex(t *testing.T) {
	var query = New()
	var date = time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local)
	if query.IsHoliday(date, Any) {
		t.Fatalf("did not expect to find a holiday at %v", date.String())
	}

	tmpFile, err := ioutil.TempFile("", "holiday_test")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	err = json.NewEncoder(tmpFile).Encode(File{Holidays: []Entry{{
		Date:   "2000-01-01",
		Name:   "test",
		States: []string{Bavaria},
	}}})

	if err != nil {
		t.Fatalf("cannot encode test entry to file: %v", err)
	}

	err = query.AddFileToIndex(tmpFile.Name())
	if err != nil {
		t.Fatalf("error adding file top index: %v", err)
	}

	if !query.IsHoliday(date, Any) {
		t.Fatalf("did expect to find a holiday at %v", date.String())
	}
}

func TestQuery_Find(t *testing.T) {
	var query = New()
	var date = time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)

	if res, ok := query.Find(date, Any); ok {
		want := "Neujahr"
		if res.Name != want {
			t.Fatalf("did expect to find '%v', but found '%v'", want, res.Name)
		}
	} else {
		t.Fatalf("did expect to find a holiday at %v", date.String())
	}
}

func parseDate(t *testing.T, input string) time.Time {
	date, err := time.Parse("2006-01-02", input)
	if err != nil {
		t.Fatalf("cannot parse date: %v", err)
	}

	return date
}
