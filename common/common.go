package common

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type AppConfigProperties map[string]string

var ConfInfo AppConfigProperties

func init() {

	profile := "local"
	if len(os.Getenv("PROFILE")) > 0 {
		profile = os.Getenv("PROFILE")
	}

	if profile == "local" {
		_, err := ReadPropertiesFile("config.properties")
		if err != nil {
			return
		}
	}

}

func ReadPropertiesFile(filename string) (AppConfigProperties, error) {
	ConfInfo = AppConfigProperties{}

	if len(filename) == 0 {
		return ConfInfo, nil
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				ConfInfo[key] = value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return ConfInfo, nil
}
