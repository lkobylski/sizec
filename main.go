package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	BYTE     = 1.0 << (10 * iota)
	KILOBYTE
	MEGABYTE
	GIGABYTE
	TERABYTE
)

func getDetectedSize(bytes int64) string {
	unit := ""
	value := float32(bytes)

	switch {

	case bytes >= TERABYTE:
		unit = "TB"
		value = value / TERABYTE
	case bytes >= GIGABYTE:
		unit = "GB"
		value = value / GIGABYTE
	case bytes >= MEGABYTE:
		unit = "MB"
		value = value / MEGABYTE
	case bytes >= KILOBYTE:
		unit = "KB"
		value = value / KILOBYTE
	case bytes == 0:
		return "0"

	}

	result := fmt.Sprintf("%.2f", value)
	result = strings.TrimSuffix(result, ".00")
	return fmt.Sprintf("%s%s", result, unit)
}

func main() {

	if len(os.Args) < 2 {
		printUsage()
	}

	bytes, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		printUsage()
	}

	fmt.Println(getDetectedSize(bytes))
	os.Exit(1)
}

func printUsage() {
	fmt.Printf("Usage: %s [bytes]\n", os.Args[0])
	fmt.Print("bytes must be integer\n", )
	os.Exit(1)
}
