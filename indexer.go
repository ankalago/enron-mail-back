package main

import (
	"chi-example/utils"
)

func main() {
	baseUrl := "./emails/enron_mail_20110402/maildir"

	//result := RecursiveScan(baseUrl, "", 0)
	//file, err := json.Marshal(result)
	//if err != nil {
	//	panic(err)
	//}
	//_ = os.WriteFile("enron_mail_20110402.json", file, 0644)
	//fmt.Println(string(u))

	result := utils.ScanDirectory(baseUrl)
	chunks := utils.ChunkSlice(result, len(result)/200)
	utils.ScanAndBulkFiles(chunks)
}
