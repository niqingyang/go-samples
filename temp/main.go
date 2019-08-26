package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {

	files := make([]string, 3)

	files[0] = "D:/Users/桌面/68patch/00 正式站/4.2/hooks.log"
	files[1] = "D:/Users/桌面/68patch/00 正式站/4.2/hooks.log.2019-08-15.1"
	files[2] = "D:/Users/桌面/68patch/00 正式站/4.2/hooks.log.2019-08-18.1"

	userIds := make(map[string]int, 1000)

	for _, filename := range files {
		if data, err := ioutil.ReadFile(filename); err == nil {
			for _, line := range strings.Split(string(data), "\n") {
				handleLine(&line, userIds)
			}
		} else {
			fmt.Printf("%v\n", err)
		}
	}

	ids := make([]string, 0, 500)

	for id, _ := range userIds {
		ids = append(ids, id)
	}

	fmt.Printf("%v\n", ids)

	data := []byte(strings.Join(ids, ","))

	ioutil.WriteFile("D:/ids.txt", data, 0777)
}

func handleLine(line *string, userIds map[string]int) {
	*line = strings.TrimSpace(*line)

	r := regexp.MustCompile("\"mobile\": \"\\d+\"")

	ids := r.FindAllString(*line, -1)

	if len(ids) > 0 {

		for _, id := range ids {
			userIds[id[11 : len(id)-1]]++
		}

	}

	//if strings.HasPrefix(*line, "")
}
