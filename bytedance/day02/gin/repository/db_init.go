package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

var (
	topicIndexMap map[int64]*Topic
)

func initTopicIndexMap(filePath string) error {
	open, err := os.Open(filePath + "topic")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	topicTmpMap := make(map[int64]*Topic)
	for scanner.Scan() {
		text := scanner.Text()
		var topic Topic
		if err := json.Unmarshal([]byte(text), &topic); err != nil {
			return err
		}
		topicTmpMap[topic.Id] = &topic
	}
	topicIndexMap = topicTmpMap
	return nil
}

func main() {
	err := initTopicIndexMap("E:\\go\\src\\bytedance\\day02\\gin\\data\\")
	if err != nil {
		fmt.Println("出现错误", err)
	}
	fmt.Println("topicIndexMap", topicIndexMap)
}
