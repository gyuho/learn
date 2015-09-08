package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fpath := "test.tsv"
	if err := toTSV([][]string{{"A", "B", "C"}, {"D", "E", "F"}}, fpath); err != nil {
		panic(err)
	}
	rows, err := fromTSV(fpath)
	if err != nil {
		panic(err)
	}
	if len(rows) != 2 {
		panic(err)
	}
	os.Remove(fpath)
}

func toTSV(rows [][]string, fpath string) error {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		f, err = os.Create(fpath)
		if err != nil {
			return err
		}
	}
	defer f.Close()

	// func NewWriter(w io.Writer) *Writer
	wr := bufio.NewWriter(f)
	for _, row := range rows {
		for idx, elem := range row {

			// func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
			fmt.Fprint(wr, elem)

			if len(row)-1 != idx {
				fmt.Fprint(wr, "\t")
			}
		}
		fmt.Fprint(wr, "\n")
	}
	if err := wr.Flush(); err != nil {
		return err
	}
	return nil
}

func fromTSV(fpath string) ([][]string, error) {
	f, err := os.OpenFile(fpath, os.O_RDONLY, 0444)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	rows := [][]string{}

	// func NewScanner(r io.Reader) *Scanner
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		rows = append(rows, strings.Split(scanner.Text(), "\t"))
	}
	if err := scanner.Err(); err != nil {
		return rows, err
	}
	return rows, f.Close()
}
