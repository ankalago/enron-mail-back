package utils

import (
	"log"
	"os"
)

type Folder struct {
	Name string
	Type bool
}

func GetFolder(url string) []Folder {
	file, err := os.Open(url)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	folder, err := file.ReadDir(0)
	if err != nil {
		log.Fatal(err)
	}
	folders := make([]Folder, len(folder))
	for i, v := range folder {
		folders[i] = Folder{
			Name: v.Name(),
			Type: v.IsDir(),
		}
	}
	return folders
}
