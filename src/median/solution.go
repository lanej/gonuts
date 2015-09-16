package median

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func solve(r *os.File, w *os.File) {
	reader := bufio.NewReader(r)

	num_s, _ := reader.ReadString('\n')
	w.WriteString(num_s)
	num, err := strconv.ParseInt(strings.Trim(num_s, "\n"), 10, 64)

	if nil != err {
		panic(err)
	}

	w.WriteString("something\n")
	w.WriteString(strconv.FormatInt(num, 10))
	w.WriteString("\n")
	w.WriteString("something else\n")
}

func main() {
	solve(os.Stdin, os.Stdout)
}
