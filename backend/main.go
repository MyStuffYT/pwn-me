package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func csvToJSON(csvFile string, keys ...string) (map[string][]string, error) {
	file, err := os.Open(csvFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	result := make(map[string][]string)
	for _, key := range keys {
		result[key] = []string{}
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, ",")
		if len(values) != len(keys) {
			return nil, fmt.Errorf("mismatch between keys and values in line: %s", line)
		}
		for i, key := range keys {
			result[key] = append(result[key], values[i])
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func main() {
	fmt.Println("Pwn-Me | BACKEND")

	fmt.Println("checking for databases...")
	if _, err := os.Stat("databases"); os.IsNotExist(err) {
		fmt.Println("no databases folder found, creating...")
		err := os.Mkdir("databases", 0755)
		if err != nil {
			fmt.Println("Error creating databases folder:", err)
			return
		}
	} else {
		fmt.Println("databases folder exists")
	}

	// Example usage of csvToJSON
	files, err := os.ReadDir("databases")
	if err != nil {
		fmt.Println("Error reading databases folder:", err)
		return
	}

	fmt.Println("found databases:")
	for _, file := range files {
		fmt.Println(file.Name())
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".csv") {
			fmt.Println("Database found: ", file.Name())
			fmt.Printf("Enter column names for %s (separated by ':'): ", file.Name())
			var input string
			fmt.Scanln(&input)
			keys := strings.Split(input, ":")
			
			jsonData, err := csvToJSON("databases/"+file.Name(), keys...)
			if err != nil {
				fmt.Println("error processing db: ", err)
				continue
			}

			outputFile := "dbpwn.json"
			out, err := os.Create(outputFile)
			if err != nil {
				fmt.Println("Error creating output file:", err)
				continue
			}
			defer out.Close()

			encoder := json.NewEncoder(out)
			encoder.SetIndent("", "  ")
			if err := encoder.Encode(jsonData); err != nil {
				fmt.Println("Error writing JSON to file:", err)
				continue
			}

			fmt.Println("Database saved to", outputFile)
		}

}
