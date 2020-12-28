package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func readFromConfigurationFile(s string) string {

	var txtlines []string
	var result string

	file, err := os.Open("configuration.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	for _, eachline := range txtlines {
		if strings.Contains(eachline, s) {
			eachline = strings.TrimPrefix(eachline, s)
			eachline = strings.TrimSpace(eachline)
			result = eachline
		}
	}

	return result
}
