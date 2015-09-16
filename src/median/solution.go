package median

import (
	"os"
)

func solve(r *os.File, w *os.File) {
	content := make([]byte, 1024)

	r.Read(content)

	w.WriteString("something\n")
	w.WriteString(string(content[:len(content)]))
	w.WriteString("something else\n")
}

func main() {
	solve(os.Stdin, os.Stdout)
}
