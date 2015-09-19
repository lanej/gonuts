package median

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func nextNumber(r *bufio.Reader) float64 {
	num_s, _ := r.ReadString('\n')
	num, err := strconv.ParseFloat(strings.Trim(num_s, "\n"), 64)

	if nil != err {
		panic(err)
	}

	return num
}

func median(s *[]float64) float64 {
	sequence := *s
	size := len(*s)
	middle := (size / 2)

	if size%2 == 0 {
		low_middle := sequence[middle-1]
		high_middle := sequence[middle]

		return (high_middle + low_middle) / 2.0
	} else {

		return sequence[middle]
	}
}

func solve(r *os.File, w *os.File) {
	reader := bufio.NewReader(r)

	count := int(nextNumber(reader))
	var sequence []float64

	for i := 0; i < count; i++ {
		sequence = append(sequence, nextNumber(reader))
		sort.Float64s(sequence)
		w.WriteString(strconv.FormatFloat(median(&sequence), 'f', 1, 32))

		if (i + 1) < count {
			w.WriteString("\n")
		}
	}
}

func main() {
	solve(os.Stdin, os.Stdout)
}
