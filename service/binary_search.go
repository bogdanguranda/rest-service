package service

import (
	"fmt"

	"github.com/bogdanguranda/rest-service/db"
	"github.com/bogdanguranda/rest-service/util/logging"
)

type BinarySearch struct {
	db     db.DB
	logger logging.Logger
}

func NewBinarySearch(db db.DB, logger logging.Logger) *BinarySearch {
	return &BinarySearch{db: db, logger: logger}
}

func (bs *BinarySearch) Search(value int) (int, int) {
	input, _ := bs.db.GetInput()

	low, high := 0, len(input)-1

	for low <= high {
		mid := (low + high) / 2
		if input[mid] == value {
			return mid, input[mid]
		} else if input[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	bs.logger.Log(logging.LogLevelDebug, "Didn't find exact value, looking for closest value within 10% range...")
	for i := 0; i < len(input); i++ {
		if float64(value)*0.9 <= float64(input[i]) && float64(input[i]) <= float64(value)*1.1 {
			bs.logger.Log(logging.LogLevelDebug, fmt.Sprintf("Found closest index for value: %d", input[i]))
			return i, input[i]
		}
	}

	return -1, -1
}
