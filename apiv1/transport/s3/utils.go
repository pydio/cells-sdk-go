package s3

import (
	"fmt"
)

const (
	// uploadPartsSteps represents the increment in MB when computing one part's size.
	// Due to Cells Server implementation, it *must* be a multiple of 10MB.
	uploadPartsSteps = int64(10 * 1024 * 1024)
)

func ComputePartSize(fileSize, defaultPartSize, maxNumberOfParts int64) (partSize int64, er error) {
	partSize = defaultPartSize * (1024 * 1024)
	if partSize%uploadPartsSteps != 0 {
		return 0, fmt.Errorf("PartSize must be a multiple of 10MB")
	}

	if int64(float64(fileSize)/float64(partSize)) < maxNumberOfParts {
		return
	}
	partSize = int64(float64(fileSize) / float64(maxNumberOfParts))
	partSize = partSize - partSize%uploadPartsSteps + uploadPartsSteps
	return
}
