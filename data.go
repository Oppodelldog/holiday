package holiday

import (
	"embed"
	"encoding/json"
	"fmt"
	"path"
)

const dataDir = "data"
const dateLayout = "2006-01-02"

//go:embed data
var data embed.FS

func mustLoadBuiltInData() IndexByDate {
	i, err := loadData()
	if err != nil {
		panic(err)
	}

	return i
}

func loadData() (IndexByDate, error) {
	dir, err := data.ReadDir(dataDir)
	if err != nil {
		return nil, fmt.Errorf("error reading data dir '%s': %w", dataDir, err)
	}

	var index = IndexByDate{}

	for _, entry := range dir {
		var fullPath = path.Join(dataDir, entry.Name())
		fileBytes, err := data.ReadFile(fullPath)
		if err != nil {
			return nil, fmt.Errorf("error reading data File '%s': %w", fullPath, err)
		}

		var data File

		err = json.Unmarshal(fileBytes, &data)
		if err != nil {
			return nil, fmt.Errorf("error decoding json data File '%s': %w", fullPath, err)
		}

		for i := range data.Holidays {
			index[data.Holidays[i].Date] = data.Holidays[i]
		}
	}

	return index, nil
}

func addToIndex(file []byte, index IndexByDate) error {
	var data File

	var err = json.Unmarshal(file, &data)
	if err != nil {
		return fmt.Errorf("error decoding json data: %w", err)
	}

	for i := range data.Holidays {
		index[data.Holidays[i].Date] = data.Holidays[i]
	}

	return nil
}
