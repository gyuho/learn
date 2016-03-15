package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	fpath := "test.csv"
	if err := toCSV([]string{"col1", "col2", "col3"}, [][]string{{"A", "B", "C"}, {"D", "E", "F"}}, fpath); err != nil {
		panic(err)
	}
	rows, err := fromCSV(fpath)
	if err != nil {
		panic(err)
	}
	if len(rows) != 3 {
		log.Fatal("must be 3 rows")
	}
	os.Remove(fpath)
}

func toCSV(header []string, rows [][]string, fpath string) error {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		f, err = os.Create(fpath)
		if err != nil {
			return err
		}
	}
	defer f.Close()

	// func NewWriter(w io.Writer) *Writer
	wr := csv.NewWriter(f)

	if err := wr.Write(header); err != nil {
		return err
	}

	if err := wr.WriteAll(rows); err != nil {
		return err
	}

	wr.Flush()
	return wr.Error()
}

func fromCSV(fpath string) ([][]string, error) {
	f, err := os.OpenFile(fpath, os.O_RDONLY, 0444)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// func NewReader(r io.Reader) *Reader
	rd := csv.NewReader(f)
	// Reading does not require `Flush`

	// in case that rows have different number of fields
	rd.FieldsPerRecord = -1

	// rd.TrailingComma = true
	// rd.TrimLeadingSpace = true
	// rd.LazyQuotes = true

	rows, err := rd.ReadAll()
	if err != nil {
		return rows, err
	}

	return rows, nil
}
