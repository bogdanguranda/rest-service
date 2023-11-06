package db

import (
	"bufio"
	"os"
	"strconv"
)

type FileDB struct {
	filename string
	input    []int
}

func NewFileDB(filename string) *FileDB {
	return &FileDB{filename: filename}
}

func (fd FileDB) GetInput() ([]int, error) {
	if len(fd.input) == 0 {
		inputData, err := fd.readInputFromFile()
		if err != nil {
			return nil, err
		}
		fd.input = inputData
	}

	return fd.input, nil
}

func (fd FileDB) readInputFromFile() ([]int, error) {
	file, err := os.Open(fd.filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers []int
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numbers, nil
}
