package median

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestSolve(t *testing.T) {
	dir := os.TempDir()
	input_file, _ := ioutil.TempFile(dir, "median")

	input_file.WriteString("10\n1\n2\n3\n4\n5\n6\n7\n8\n9\n10")
	input_file.Seek(0, 0)

	output_file, _ := ioutil.TempFile(dir, "median")

	solve(input_file, output_file)

	output_bytes, _ := ioutil.ReadFile(output_file.Name())

	t.Logf(string(output_bytes[:len(output_bytes)]))
	t.Logf("starts")
}
