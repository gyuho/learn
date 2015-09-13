package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	func() {
		if err := fromLines([]string{"A", "B", "C"}, "./tmp.txt"); err != nil {
			panic(err)
		}
		defer os.RemoveAll("./tmp.txt")
		lines, err := toLines1("./tmp.txt")
		if err != nil {
			panic(err)
		}
		if len(lines) != 3 {
			panic(fmt.Errorf("expected 3 but %v", lines))
		}
	}()

	func() {
		if err := fromLines([]string{"A", "B", "C"}, "./tmp.txt"); err != nil {
			panic(err)
		}
		defer os.RemoveAll("./tmp.txt")
		lines, err := toLines2("./tmp.txt")
		if err != nil {
			panic(err)
		}
		if len(lines) != 3 {
			panic(fmt.Errorf("expected 3 but %v", lines))
		}
	}()

	func() {
		if err := fromLines([]string{"aaa bbb ccc"}, "./tmp.txt"); err != nil {
			panic(err)
		}
		defer os.RemoveAll("./tmp.txt")
		words, err := toWords("./tmp.txt")
		if err != nil {
			panic(err)
		}
		if len(words) != 3 {
			panic(fmt.Errorf("expected 3 but %v", words))
		}
	}()

	func() {
		fpath := "stdout.txt"
		d, err := toBytes(fpath)
		if err != nil {
			panic(err)
		}
		fmt.Println("toBytes:", string(d))
		/*
		   toBytes: Enter text: input:fmt.Println
		   Enter text: input:fmt.Println
		   Hello World!
		*/
	}()
}

func fromLines(lines []string, fpath string) error {
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

	for _, line := range lines {
		// func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
		fmt.Fprintln(wr, line)
	}

	if err := wr.Flush(); err != nil {
		return err
	}
	return nil
}

func toLines1(fpath string) ([]string, error) {
	f, err := os.OpenFile(fpath, os.O_RDONLY, 0444)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines := []string{}

	// func NewScanner(r io.Reader) *Scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return lines, err
	}
	return lines, nil
}

func toLines2(fpath string) ([]string, error) {
	f, err := os.OpenFile(fpath, os.O_RDONLY, 0444)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	rs := []string{}
	br := bufio.NewReader(f)
	for {
		l, err := br.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		rs = append(rs, l)
	}
	return rs, nil
}

func toWords(fpath string) ([]string, error) {
	f, err := os.OpenFile(fpath, os.O_RDONLY, 0444)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines := []string{}

	// func NewScanner(r io.Reader) *Scanner
	scanner := bufio.NewScanner(f)

	// This must be called before Scan.
	// The default split function is bufio.ScanLines.
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return lines, err
	}
	return lines, nil
}

func toBytes(fpath string) ([]byte, error) {
	f, err := os.OpenFile(fpath, os.O_RDONLY, 0444)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	rs := []byte{}
	br := bufio.NewReader(f)
	for {
		c, err := br.ReadByte()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		rs = append(rs, c)
	}
	return rs, nil
}
