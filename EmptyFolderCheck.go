package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	var rootDir string
	if len(os.Args) > 1 {
		rootDir = os.Args[1]
	} else {
		var err error
		rootDir, err = os.Getwd()
		if err != nil {
			fmt.Println("Failed to get current directory:", err)
			return
		}
	}

	emptyDirs := findEmptyDirs(rootDir)
	if len(emptyDirs) == 0 {
		fmt.Println("No empty directories found.")
		return
	}

	for i, dir := range emptyDirs {
		fmt.Printf("%d: %s\n", i+1, dir)
	}

	fmt.Println("Enter the numbers of directories to delete (e.g. 1,2,4-7):")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	toDelete := parseInput(input)
	for _, idx := range toDelete {
		if idx >= 1 && idx <= len(emptyDirs) {
			err := os.Remove(emptyDirs[idx-1])
			if err != nil {
				fmt.Printf("Failed to delete %s: %s\n", emptyDirs[idx-1], err)
			} else {
				fmt.Printf("Deleted %s\n", emptyDirs[idx-1])
			}
		} else {
			fmt.Printf("Invalid index: %d\n", idx)
		}
	}
}

func findEmptyDirs(root string) []string {
	var emptyDirs []string
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			entries, err := os.ReadDir(path)
			if err != nil {
				return err
			}
			if len(entries) == 0 {
				emptyDirs = append(emptyDirs, path)
			}
		}
		return nil
	})
	return emptyDirs
}

func parseInput(input string) []int {
	var result []int
	parts := strings.Split(input, ",")
	for _, part := range parts {
		if strings.Contains(part, "-") {
			bounds := strings.Split(part, "-")
			if len(bounds) == 2 {
				start, err1 := strconv.Atoi(bounds[0])
				end, err2 := strconv.Atoi(bounds[1])
				if err1 == nil && err2 == nil && start <= end {
					for i := start; i <= end; i++ {
						result = append(result, i)
					}
				}
			}
		} else {
			num, err := strconv.Atoi(part)
			if err == nil {
				result = append(result, num)
			}
		}
	}
	return result
}
