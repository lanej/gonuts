package median

import (
	"bufio"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	input := "10\n1\n2\n3\n4\n5\n6\n7\n8\n9\n10"
	reader := strings.NewReader(input)
	solve(bufio.NewReader(reader))
}
