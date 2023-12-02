package shared

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

type DayChallenge struct {
	DayNum int
	Path   string
	Inputs []string
}

func NewDayChallenge(dayNum int) (*DayChallenge, error) {
	dayChallenge := DayChallenge{
		DayNum: dayNum,
		Path:   fmt.Sprintf("./day%d/input.txt", dayNum),
	}

	err := dayChallenge.getDataInput()
	if err != nil {
		return nil, err
	}

	return &dayChallenge, nil
}

func (d *DayChallenge) getDataInput() error {
	file, err := os.Open(d.Path)
	if errors.Is(err, os.ErrNotExist) {
		// handle the case where the file doesn't exist
		log.Println(d.Path)
		return fmt.Errorf("input file from day %d doesn't exist", d.DayNum)
	}
	if err != nil {
		return err
	}

	err = d.readLines(file)
	if err != nil {
		return err
	}

	return nil
}

func (d *DayChallenge) readLines(f *os.File) error {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		d.Inputs = append(d.Inputs, scanner.Text())
	}

	if err := scanner.Err(); err != nil {

		return err
	}

	return nil
}
