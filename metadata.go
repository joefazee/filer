package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Metadata struct {
	OriginalFileName string `json:"original_file_name"`
	PartSize         int64  `json:"part_size"`
	TotalParts       int64  `json:"total_parts"`
}

func SaveMetadata(filePath string, partSize, totalParts int64) error {

	metaData := Metadata{
		OriginalFileName: filepath.Base(filePath),
		PartSize:         partSize,
		TotalParts:       totalParts,
	}

	mbs, err := json.Marshal(metaData)
	if err != nil {
		return err
	}

	mdf := filePath + ".metadata.json"
	if err := os.WriteFile(mdf, mbs, 0644); err != nil {
		return err
	}

	return nil

}

func LoadMetadata(filePath string) (*Metadata, error) {
	mdf := filePath + ".metadata.json"

	mbs, err := os.ReadFile(mdf)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("%s does not exist", mdf)
		}
		return nil, err
	}

	var mdata Metadata
	if err := json.Unmarshal(mbs, &mdata); err != nil {
		return nil, err
	}

	return &mdata, nil
}
