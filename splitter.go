package main

import (
	"fmt"
	"io"
	"os"
)

func SplitFile(filePath string, partSize int64) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("error getting file inf %v", err)
	}

	totalParts := fileInfo.Size() / partSize
	if fileInfo.Size()%partSize != 0 {
		totalParts++
	}

	for i := int64(0); i < totalParts; i++ {
		unProcessedBs := min(partSize, fileInfo.Size()-i*partSize)
		outFile, err := os.Create(fmt.Sprintf("%s.part_%d", filePath, i+1))
		if err != nil {
			return fmt.Errorf("error creating part file: %v", err)
		}
		_, err = io.CopyN(outFile, file, unProcessedBs)
		if err != nil {
			return fmt.Errorf("error writing file part: %v", err)
		}

		if err = outFile.Close(); err != nil {
			return fmt.Errorf("error closing file %v", err)
		}
	}

	if err := SaveMetadata(filePath, partSize, totalParts); err != nil {
		return fmt.Errorf("error writing metadata: %v", err)
	}

	_ = os.Remove(filePath)

	return nil
}

func JoinFiles(filePath string) error {
	mdata, err := LoadMetadata(filePath)
	if err != nil {
		return err
	}

	outFile, err := os.Create(fmt.Sprintf("%s", mdata.OriginalFileName))
	if err != nil {
		return err
	}
	defer outFile.Close()

	for i := int64(0); i < mdata.TotalParts; i++ {
		partFileName := fmt.Sprintf("%s.part_%d", filePath, i+1)
		partFile, err := os.Open(partFileName)
		if err != nil {
			return err
		}
		_, err = io.Copy(outFile, partFile)
		if err != nil {
			return err
		}
		_ = os.Remove(partFileName)
		if err = partFile.Close(); err != nil {
			return fmt.Errorf("error closing file %v", err)
		}
	}

	_ = os.Remove(filePath + ".metadata.json")
	return nil
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
