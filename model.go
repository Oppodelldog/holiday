package holiday

// Entry defines one holiday date.
type Entry struct {
	Date   string `json:"date"`
	Name   string `json:"name"`
	States States `json:"states,omitempty"`
}

// States holds those states the holiday date is applied to.
type States []string

// File wraps holiday date entries for file serialization.
type File struct {
	Holidays []Entry `json:"holidays"`
}

// IndexByDate holds Entry indexed by Entry.Date.
type IndexByDate map[string]Entry
