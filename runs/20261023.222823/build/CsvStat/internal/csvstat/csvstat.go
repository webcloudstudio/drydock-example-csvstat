package csvstat

import (
	"encoding/csv"
	"fmt"
	"io"
	"math/big"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Record []string
type Input struct {
	Header  Record
	Records []Record
}

func Load(reader io.Reader) (Input, error) {
	r := csv.NewReader(reader)
	r.FieldsPerRecord = -1
	rows, err := r.ReadAll()
	if err != nil {
		return Input{}, fmt.Errorf("read csv: %w", err)
	}
	if len(rows) == 0 {
		return Input{}, fmt.Errorf("empty input")
	}
	header := trimRecord(rows[0])
	if len(header) == 0 {
		return Input{}, fmt.Errorf("empty header")
	}
	seen := map[string]bool{}
	for _, name := range header {
		if name == "" || seen[name] {
			return Input{}, fmt.Errorf("invalid or repeated header")
		}
		seen[name] = true
	}
	out := Input{Header: header}
	for i, row := range rows[1:] {
		if len(row) != len(header) {
			return Input{}, fmt.Errorf("record line %d has %d fields, want %d", i+2, len(row), len(header))
		}
		out.Records = append(out.Records, trimRecord(row))
	}
	return out, nil
}
func trimRecord(record []string) Record {
	out := make(Record, len(record))
	for i, v := range record {
		out[i] = strings.TrimSpace(v)
	}
	return out
}
func IsNull(v string) bool {
	v = strings.ToUpper(v)
	return v == "" || v == "NA" || v == "N/A" || v == "NULL"
}

type ProfileRow struct{ Column, Type, Rows, Nulls, Distinct, Min, Max, Sum, Mean, Median string }

func (r ProfileRow) CSV() []string {
	return []string{r.Column, r.Type, r.Rows, r.Nulls, r.Distinct, r.Min, r.Max, r.Sum, r.Mean, r.Median}
}

func Profile(in Input, selected string) ([]ProfileRow, error) {
	indices, err := columnIndices(in.Header, selected)
	if err != nil {
		return nil, err
	}
	out := make([]ProfileRow, 0, len(indices))
	for _, ix := range indices {
		values := []string{}
		nulls := 0
		for _, rec := range in.Records {
			if IsNull(rec[ix]) {
				nulls++
			} else {
				values = append(values, rec[ix])
			}
		}
		row := ProfileRow{Column: in.Header[ix], Rows: strconv.Itoa(len(in.Records)), Nulls: strconv.Itoa(nulls), Distinct: strconv.Itoa(distinct(values)), Type: inferType(values)}
		if len(values) > 0 {
			row.Min, row.Max = extrema(values, row.Type)
			if row.Type == "integer" || row.Type == "number" {
				row.Sum, row.Mean, row.Median = numericStats(values, row.Type)
			}
		}
		out = append(out, row)
	}
	return out, nil
}
func columnIndices(header Record, selected string) ([]int, error) {
	if selected == "" {
		r := make([]int, len(header))
		for i := range r {
			r[i] = i
		}
		return r, nil
	}
	for i, n := range header {
		if n == selected {
			return []int{i}, nil
		}
	}
	return nil, fmt.Errorf("column %q not found", selected)
}
func distinct(values []string) int {
	m := map[string]bool{}
	for _, v := range values {
		m[v] = true
	}
	return len(m)
}

func inferType(values []string) string {
	if len(values) == 0 {
		return "empty"
	}
	integer, number, date, boolean := true, true, true, true
	for _, v := range values {
		if !isInteger(v) {
			integer = false
		}
		if !isNumber(v) {
			number = false
		}
		if !isDate(v) {
			date = false
		}
		if !strings.EqualFold(v, "true") && !strings.EqualFold(v, "false") {
			boolean = false
		}
	}
	if integer {
		return "integer"
	}
	if number {
		return "number"
	}
	if date {
		return "date"
	}
	if boolean {
		return "boolean"
	}
	return "text"
}
func isInteger(s string) bool {
	if s == "" {
		return false
	}
	if s[0] == '-' {
		s = s[1:]
	}
	if s == "" {
		return false
	}
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
func isNumber(s string) bool {
	if s == "" {
		return false
	}
	if s[0] == '-' {
		s = s[1:]
	}
	p := strings.Split(s, ".")
	if len(p) > 2 || p[0] == "" {
		return false
	}
	for _, part := range p {
		if part == "" {
			return false
		}
		for _, c := range part {
			if c < '0' || c > '9' {
				return false
			}
		}
	}
	return true
}
func isDate(s string) bool {
	if len(s) != 10 || s[4] != '-' || s[7] != '-' {
		return false
	}
	_, err := time.Parse("2006-01-02", s)
	return err == nil
}
func parseRat(s string) *big.Rat {
	r, ok := new(big.Rat).SetString(s)
	if !ok {
		return new(big.Rat)
	}
	return r
}
func extrema(values []string, typ string) (string, string) {
	if typ == "boolean" {
		for _, v := range values {
			if strings.EqualFold(v, "false") {
				return "false", "true"
			}
		}
		return "true", "true"
	}
	a := append([]string(nil), values...)
	if typ == "integer" || typ == "number" {
		sort.Slice(a, func(i, j int) bool { return parseRat(a[i]).Cmp(parseRat(a[j])) < 0 })
	} else {
		sort.Strings(a)
	}
	return a[0], a[len(a)-1]
}

func numericStats(values []string, typ string) (string, string, string) {
	sum := new(big.Rat)
	nums := make([]*big.Rat, len(values))
	for i, v := range values {
		nums[i] = parseRat(v)
		sum.Add(sum, nums[i])
	}
	sort.Slice(nums, func(i, j int) bool { return nums[i].Cmp(nums[j]) < 0 })
	mean := new(big.Rat).Quo(sum, big.NewRat(int64(len(nums)), 1))
	median := nums[len(nums)/2]
	if len(nums)%2 == 0 {
		median = new(big.Rat).Quo(new(big.Rat).Add(nums[len(nums)/2-1], nums[len(nums)/2]), big.NewRat(2, 1))
	}
	if typ == "integer" {
		return sum.Num().String(), formatRat(mean), formatRat(median)
	}
	return formatRat(sum), formatRat(mean), formatRat(median)
}
func formatRat(r *big.Rat) string {
	scaled := new(big.Int).Mul(r.Num(), big.NewInt(100))
	q, rem := new(big.Int), new(big.Int)
	q.QuoRem(scaled, r.Denom(), rem)
	if rem.Sign() != 0 && new(big.Int).Mul(new(big.Int).Abs(rem), big.NewInt(2)).Cmp(r.Denom()) >= 0 {
		if r.Sign() >= 0 {
			q.Add(q, big.NewInt(1))
		} else {
			q.Sub(q, big.NewInt(1))
		}
	}
	neg := q.Sign() < 0
	s := new(big.Int).Abs(q).String()
	for len(s) < 3 {
		s = "0" + s
	}
	s = s[:len(s)-2] + "." + s[len(s)-2:]
	if neg {
		s = "-" + s
	}
	return s
}

type Frequency struct {
	Value string
	Count int
}

func Frequencies(in Input, selected string, top int) ([]Frequency, error) {
	ix, err := columnIndices(in.Header, selected)
	if err != nil {
		return nil, err
	}
	counts := map[string]int{}
	for _, r := range in.Records {
		if !IsNull(r[ix[0]]) {
			counts[r[ix[0]]]++
		}
	}
	out := make([]Frequency, 0, len(counts))
	for v, c := range counts {
		out = append(out, Frequency{v, c})
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Count != out[j].Count {
			return out[i].Count > out[j].Count
		}
		return out[i].Value < out[j].Value
	})
	if len(out) > top {
		out = out[:top]
	}
	return out, nil
}
func RenderProfile(w io.Writer, rows []ProfileRow) error {
	c := csv.NewWriter(w)
	if err := c.Write([]string{"column", "type", "rows", "nulls", "distinct", "min", "max", "sum", "mean", "median"}); err != nil {
		return err
	}
	for _, r := range rows {
		if err := c.Write(r.CSV()); err != nil {
			return err
		}
	}
	c.Flush()
	return c.Error()
}
func RenderFrequency(w io.Writer, values []Frequency) error {
	c := csv.NewWriter(w)
	if err := c.Write([]string{"value", "count"}); err != nil {
		return err
	}
	for _, v := range values {
		if err := c.Write([]string{v.Value, strconv.Itoa(v.Count)}); err != nil {
			return err
		}
	}
	c.Flush()
	return c.Error()
}
