// Package pho provides pho  
package pho

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"slices"
)

// Analyses struct  
type Analyses struct {
	histogram map[int]int
	title     string
}

// Data method  
func (a *Analyses) Data() map[int]int {
	return a.histogram
}

// Title method  
func (a *Analyses) Title() string {
	return a.title
}

// Analyze function  
func Analyze(dirpaths []string, recurse bool) ([]*Analyses, error) {
	analyses := make([]*Analyses, 0)
	for _, dir := range dirpaths {
		files, err := listFiles(dir, recurse)
		filteredFiles := make([]string, 0)
		if err != nil {
			return nil, fmt.Errorf("failed to list files in directory %s: %w", dir, err)
		}
		for _, file := range files {
			basename := path.Base(file)
			isBlackListed := !slices.Contains([]string{"box.html", "go.mod", "go.sum"}, basename)
			fmt.Println(basename, isBlackListed)
			if !hiddenPathFilter(file) &&
				isBlackListed {
				filteredFiles = append(filteredFiles, file)
			}
		}
		fmt.Println(filteredFiles)
		lines, err := readLinesFromFiles(filteredFiles)
		if err != nil {
			return nil, fmt.Errorf("failed to read lines from files in directory %s: %w", dir, err)
		}
		histmap := lineLengthFrequency(lines)
		analyses = append(analyses, &Analyses{histmap, filepath.Base(dir)})
	}
	return analyses, nil
}

func lineLengthFrequency(lines []string) map[int]int {
	frequencyMap := make(map[int]int)

	for _, line := range lines {
		lineLength := len(line)
		frequencyMap[lineLength]++
	}

	return frequencyMap
}

func listFiles(dirpath string, recurse bool) ([]string, error) {
	var files []string
	dirp, err := filepath.Abs(dirpath)
	if err != nil {
		return nil, err
	}
	err = filepath.WalkDir(dirp, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			files = append(files, path)
		} else if !recurse && path != dirpath {
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, nil
}

func readLinesFromFiles(filepaths []string) ([]string, error) {
	lines := make([]string, 0)

	for _, filepath := range filepaths {
		file, err := os.Open(filepath)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}

	return lines, nil
}
