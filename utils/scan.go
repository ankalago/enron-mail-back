package utils

import (
	"chi-example/constants"
	"chi-example/entities"
	"chi-example/httpd"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

func ScanAndBulkFiles(chunks [][]string) {
	var wg sync.WaitGroup
	var limits int

	routine := func(files []string, wg *sync.WaitGroup) {
		var messages []Message
		textResult := ""
		defer wg.Done()
		for _, item := range files {
			messages = append(messages, ScanFile(item))

		}
		for _, file := range messages {
			b, _ := json.Marshal(file)
			textResult += fmt.Sprintf("%s\n%s\n", constants.IndexObject, string(b))
			limits += len(file.Content)
		}
		avg := limits / len(messages)
		httpd.BulkData(textResult, avg)
		fmt.Println("End bash block")
	}

	for _, chunk := range chunks {
		wg.Add(1)
		go routine(chunk, &wg)
	}
	wg.Wait()
}

func RecursiveScan(baseUrl string, path string, index int) []any {
	folders := GetFolder(baseUrl + path)
	//spaces := strings.Repeat(" ", index)
	var array = make([]any, len(folders))
	//var slice []Node[Items]
	for i, v := range folders {
		if v.Name == ".DS_Store" {
			//fmt.Printf("%s|-- %s \n", spaces, v.name)
			//slice = append(slice, Node[Item]{name: v.name, Items: nil})
			continue
		}
		// add file
		content := ""
		if !v.Type {
			cont, err := os.ReadFile(baseUrl + path + "/" + v.Name)
			if err != nil {
				log.Fatal(err)
			}
			content = string(cont)
		}
		array[i] = entities.File{Name: v.Name, Content: content}

		// add folder
		if v.Type {
			arrayFolder := RecursiveScan(baseUrl+path, "/"+v.Name, index+4)
			array[i] = entities.Folder{Name: v.Name, Items: arrayFolder}
			//slice = append(slice, Node[Items]{name: v.name, Items: nil})
		}
	}
	return RemoveNullValue(array)
}
