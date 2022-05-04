package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type JsonLog map[string]interface{}

func main() {
	var x, err = lineByLine()
	if err != nil {
		panic(err)
	}
	result := joinString(uniqueKeys(x)) + addValueToMapString(x, uniqueKeys(x))
	err = os.WriteFile("E:\\work\\tabl.csv", []byte(result), 0644)
	if err != nil {
		panic(err)
	}
}

func joinString(key []string) string {
	return strings.Join(key, ",") + "\n"
}

func addValueToMapString(listLog []JsonLog, keys []string) string {
	var result string
	for _, v := range listLog {
		var valueMap []string

		for _, k := range keys {
			valueMap = append(valueMap, toString(v[k]))
		}
		result += joinString(valueMap)
	}
	return result
}

func toString(i interface{}) string {
	if i == nil {
		return ""
	}
	return fmt.Sprintf("%v", i)
}

func uniqueKeys(listLog []JsonLog) []string {
	var keys []string
	keysSet := map[string]bool{}

	for logs := range listLog {
		for key := range listLog[logs] {
			if _, ok := keysSet[key]; ok {
				continue
			}
			keys = append(keys, key)
			keysSet[key] = true
		}
	}
	return keys
}

func lineByLine() ([]JsonLog, error) {
	var listLog []JsonLog

	file, err := os.Open("E:\\work\\test.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var logL1 JsonLog
		err = json.Unmarshal(scanner.Bytes(), &logL1)
		if err != nil {
			return nil, err
		}

		logLStr := logL1["log"].(string)
		var logL2 JsonLog
		err = json.Unmarshal([]byte(logLStr), &logL2)
		if err != nil {
			return nil, err
		}
		listLog = append(listLog, logL2)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return listLog, nil
}
