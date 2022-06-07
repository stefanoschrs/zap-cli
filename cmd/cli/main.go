package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var data map[string]interface{}
		err := json.Unmarshal(scanner.Bytes(), &data)
		if err != nil {
			fmt.Println(scanner.Text())
			continue
		}
		time, ok := data["time"].(string)
		if ok {
			delete(data, "time")
		}
		level, ok := data["level"].(string)
		if ok {
			delete(data, "level")
		}
		message, ok := data["message"].(string)
		if ok {
			delete(data, "message")
		}

		if len(data) > 0 {
			dataStr, err2 := json.Marshal(data)
			if err2 != nil {
				fmt.Printf("%s %s\t%s\t%+v\n", time, level, message, data)
			} else {
				fmt.Printf("%s %s\t%s\t%s\n", time, level, message, dataStr)
			}
		} else {
			fmt.Printf("%s %s\t%s\n", time, level, message)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(fmt.Errorf("scanner.Err: %w", err))
	}
}
