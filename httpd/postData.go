package httpd

import (
	"chi-example/constants"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func BulkData(data string, avg int) {
	req, err := http.NewRequest("POST", constants.Endpoint+"_bulk", strings.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(constants.User, constants.Password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)

	errWrite := os.WriteFile(fmt.Sprintf("./logs/%d-%d.json", time.Now().Unix(), resp.StatusCode), []byte(data), 0644)
	if errWrite != nil {
		log.Fatal("errWrite: ", errWrite)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body), avg)
}
