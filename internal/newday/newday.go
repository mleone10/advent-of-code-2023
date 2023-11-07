package newday

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed day.go.template
var dayTemplate string

//go:embed day_test.go.template
var dayTestTemplate string

var tDay = template.Must(template.New("day").Parse(dayTemplate))
var tDayTest = template.Must(template.New("dayTest").Parse(dayTestTemplate))

type templateValues struct {
	DayNum int
}

func Init(basePath string, dayNum int) error {
	dayName := fmt.Sprintf("day%02d", dayNum)
	dayDirPath := filepath.Join(basePath, dayName)

	err := os.Mkdir(dayDirPath, 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory at %v: %w", dayDirPath, err)
	}

	tvs := templateValues{DayNum: dayNum}
	err = writeTemplateToFile(tDay, tvs, filepath.Join(dayDirPath, fmt.Sprintf("%s.go", dayName)))
	if err != nil {
		return err
	}
	err = writeTemplateToFile(tDayTest, tvs, filepath.Join(dayDirPath, fmt.Sprintf("%s_test.go", dayName)))
	if err != nil {
		return err
	}

	_, err = createFile(filepath.Join(dayDirPath, "input.txt"))
	if err != nil {
		return err
	}

	_, err = createFile(filepath.Join(dayDirPath, "test_input.txt"))
	if err != nil {
		return err
	}

	return nil
}

func writeTemplateToFile(t *template.Template, tvs templateValues, filePath string) error {
	file, err := createFile(filePath)
	if err != nil {
		return err
	}

	err = t.Execute(file, tvs)
	if err != nil {
		return fmt.Errorf("failed to write template data to file: %w", err)
	}

	return nil
}

func createFile(path string) (*os.File, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("failed to create empty file: %w", err)
	}
	return file, nil
}

func dayNameString(dayNum int) string {
	return fmt.Sprintf("day%02d", dayNum)
}
