package utils

import (
	"bufio"
	"chi-example/constants"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Folder struct {
	Name string
	Type bool
}

type Message struct {
	To      string `json:"to"`
	From    string `json:"from"`
	Subject string `json:"subject"`
	Origin  string `json:"origin"`
	Content string `json:"content"`
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

func ScanDirectory(baseUrl string) []Message {
	var messages []Message
	err := filepath.Walk(baseUrl,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				replace := strings.Replace(path, baseUrl, "", 1)
				if replace != "" {
					//fmt.Printf("|-- %s\n", replace)
				}
			} else {
				if info.Name() != ".DS_Store" {
					//fmt.Printf("    |-- %s\n", info.name()) // path
					messages = append(messages, scanFile(path))
				}
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return messages
}

func scanFile(url string) Message {
	file, err := os.Open(url)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtLines []string

	aux := 0
	indexLineContent := 0
	indexLineOrigin := 0
	var xOrigin int
	for scanner.Scan() {
		txtLines = append(txtLines, ValidateEmptyText(scanner.Text()))
		xFileName := strings.Index(scanner.Text(), constants.XFileName)
		xOrigin = strings.Index(scanner.Text(), constants.XOrigin)
		if xFileName >= 0 {
			indexLineContent = aux
		}
		if xOrigin >= 0 {
			indexLineOrigin = aux
		}
		aux++
	}

	defer file.Close()

	message := Message{
		To:      FormatText(txtLines[3]),
		From:    FormatText(txtLines[2]),
		Subject: FormatText(txtLines[4]),
		Origin:  FormatText(strings.ToLower(txtLines[indexLineOrigin])),
		Content: strings.Join(txtLines[indexLineContent+1:], ""),
	}
	return message
}
