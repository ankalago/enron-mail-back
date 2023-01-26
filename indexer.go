package main

import (
	"chi-example/utils"
	"encoding/json"
	"os"
)

type cacheable interface {
	Items | Item
}

type Items struct {
	Name  string `json:"name"`
	Items []any  `json:"items"`
}

type Item struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type Node[T cacheable] struct {
	Name string `json:"name"`

	Items []T `json:"items"`
}

func main() {
	//baseUrl := "/Users/pauljacome/Downloads/enron_mail_20110402/test"
	baseUrl := "/Users/pauljacome/Downloads/enron_mail_20110402/maildir"
	result := recursiveScan(baseUrl, "", 0)
	file, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	_ = os.WriteFile("enron_mail_20110402.json", file, 0644)
	//fmt.Println(string(u))
}

func recursiveScan(baseUrl string, path string, index int) []any {
	folders := utils.GetFolder(baseUrl + path)
	//spaces := strings.Repeat(" ", index)
	var array = make([]any, len(folders))
	//var slice []Node[Items]
	for i, v := range folders {
		if v.Name == ".DS_Store" {
			//fmt.Printf("%s|-- %s \n", spaces, v.Name)
			//slice = append(slice, Node[Item]{Name: v.Name, Items: nil})
			continue
		}
		// add file
		content := ""
		//if !v.Type {
		//	cont, err := os.ReadFile(baseUrl + path + "/" + v.Name)
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//	content = string(cont)
		//}
		array[i] = Item{Name: v.Name, Content: content}

		// add folder
		if v.Type {
			arrayFolder := recursiveScan(baseUrl+path, "/"+v.Name, index+4)
			array[i] = Items{Name: v.Name, Items: arrayFolder}
			//slice = append(slice, Node[Items]{Name: v.Name, Items: nil})
		}
	}
	return utils.RemoveNullValue(array)
}
