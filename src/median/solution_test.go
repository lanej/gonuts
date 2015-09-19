package median

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestSolve(t *testing.T) {
	dir := os.TempDir()
	input_file, _ := ioutil.TempFile(dir, "median")

	// rewite as a go routine
	input_file.WriteString("10\n1\n2\n3\n4\n5\n6\n7\n8\n9\n10")
	input_file.Seek(0, 0)

	output_file, _ := ioutil.TempFile(dir, "median")

	solve(input_file, output_file)

	output_bytes, _ := ioutil.ReadFile(output_file.Name())

	output := string(output_bytes[:len(output_bytes)])
	expected_output := "1.0\n1.5\n2.0\n2.5\n3.0\n3.5\n4.0\n4.5\n5.0\n5.5"
	assert.Equal(t, expected_output, output)
}
