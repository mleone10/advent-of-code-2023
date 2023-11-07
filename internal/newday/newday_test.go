package newday_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mleone10/advent-of-code-2023/internal/newday"
)

const dayNum = 99

var dayName = fmt.Sprintf("day%02d", dayNum)

func TestFilesCreated(t *testing.T) {
	tmp := t.TempDir()
	tcs := []string{
		filepath.Join(tmp, fmt.Sprintf("%[1]s/%[1]s.go", dayName)),
		filepath.Join(tmp, fmt.Sprintf("%[1]s/%[1]s_test.go", dayName)),
		filepath.Join(tmp, fmt.Sprintf("%s/input.txt", dayName)),
		filepath.Join(tmp, fmt.Sprintf("%s/test_input.txt", dayName)),
	}

	err := newday.Init(tmp, dayNum)
	if err != nil {
		t.Fatalf("newday initialization failed: %v", err)
	}

	for _, tc := range tcs {
		if _, err := os.Stat(tc); errors.Is(err, os.ErrNotExist) {
			t.Errorf("file %s does not exist, but should", tc)
		}
	}
}

func TestSourceFilesPopulated(t *testing.T) {
	tmp := t.TempDir()
	tcs := []struct {
		file      string
		firstLine string
	}{
		{
			filepath.Join(tmp, fmt.Sprintf("%[1]s/%[1]s.go", dayName)),
			fmt.Sprintf("package %s", dayName),
		}, {
			filepath.Join(tmp, fmt.Sprintf("%[1]s/%[1]s_test.go", dayName)),
			fmt.Sprintf("package %s_test", dayName),
		},
	}

	err := newday.Init(tmp, dayNum)
	if err != nil {
		t.Fatalf("newday initialization failed: %v", err)
	}

	for _, tc := range tcs {
		bs, err := os.ReadFile(tc.file)
		if err != nil {
			t.Errorf("failed to open %v to read: %v", tc.file, err)
			continue
		}
		actual := strings.Split(string(bs), "\n")[0]
		if actual != tc.firstLine {
			t.Errorf("first line [%s] did not match expected [%s]", actual, tc.firstLine)
		}
	}
}
