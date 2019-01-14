package reader

import (
	"encoding/csv"
	"io/ioutil"
	"os"
	"strings"
)

func ReadFile(filepath string) (records [][]string, err error) {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return
	}
	r := csv.NewReader(strings.NewReader(string(b)))
	records, err = r.ReadAll()
	return
}

func CheckFile(filepath string) error {
	_, err := os.Stat(filepath)
	return err
}