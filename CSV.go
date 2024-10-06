package dbt1

import (
	"encoding/csv"
	"strings"
	"bytes"
)

// CSV Holds a list of values separated by commas.
type CSV string

// Contains returns true if the CSV contains the given value.
func (c CSV) Contains(s string) bool {
	for _, v := range c.Getl() {
		if v == s { return true }
	}
	return false
}

// Setl sets the CSV to the given list of values.
func (c *CSV) Setl(l []string) {
	var w *csv.Writer
	var b []byte
	w = csv.NewWriter(bytes.NewBuffer(b))
	w.Write(l)
	w.Flush()
	*c = CSV(b)
}

// Getl returns the list of values in the CSV.
func (c  CSV) Getl() (sl []string) {
	var r   *csv.Reader
	var i    int
	var err  error
	r = csv.NewReader(strings.NewReader(string(c)))
	sl = []string{}
	sl, err = r.Read()
	if err != nil && err.Error() != "EOF" { panic(err) }
	for i = range sl {
		sl[i] = strings.TrimSpace(sl[i])
	}
	return
}

// Sets sets the CSV to the given list of values. Does nothing if
// the given string is "-".
func (c *CSV) Sets(s string) {
	if (s != "-") {
		*c = CSV(s)
	}
}

func (c CSV) Get(n int) (r string) {
	var sl []string = c.Getl()
	if n < 0 || n >= len(sl) {
		return ""
	}
	return sl[n]
}

func (c CSV) Has(n int) (b bool) {
	var sl []string = c.Getl()
	return n >= 0 && n < len(sl)
}
