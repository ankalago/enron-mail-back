package httpd

import (
	"chi-example/constants"
	"chi-example/entities"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func GetData(queryText string) entities.Request {
	where := ""
	if queryText != "" {
		where = fmt.Sprintf(` WHERE \"origin\"='%s'`, queryText)
	}
	query := fmt.Sprintf(`{
			"query": {
				"sql": "SELECT * FROM \"%s\"%s",
				"start_time": 1672549200000000,
				"end_time": 1675227540000000,
				"from": 0,
				"size": 1000
			}
		}`, constants.Index, where)
	req, err := http.NewRequest("POST", "https://api.zincsearch.com/api/enterprise_16O476E9609y7Lz/_search", strings.NewReader(query))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", "Bearer eyJraWQiOiJ1TVJxZlRhUmp1c1pyWUZFRWVFMlVFeUVhQW50WVhMbEZOQU11eitIZkZFPSIsImFsZyI6IlJTMjU2In0.eyJhdF9oYXNoIjoiQzJTdjdIOTBqUVBGZTBRUjFkdmpWUSIsInN1YiI6ImQxM2NjMDA1LWNmZGEtNDM0Ny05YWU4LWQxMzIxZmM1ODAzZCIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJpc3MiOiJodHRwczpcL1wvY29nbml0by1pZHAudXMtZWFzdC0yLmFtYXpvbmF3cy5jb21cL3VzLWVhc3QtMl9weTJnSVlZOXYiLCJjb2duaXRvOnVzZXJuYW1lIjoiZDEzY2MwMDUtY2ZkYS00MzQ3LTlhZTgtZDEzMjFmYzU4MDNkIiwib3JpZ2luX2p0aSI6ImIxOTI2ODZkLTI3NzYtNDRlMy1iMzkzLTgxMWFmNGJkOWE1MSIsImF1ZCI6IjNhc2RjZHR0aWducGY1cWpzN25tdDRmNjY2IiwiZXZlbnRfaWQiOiI4ODVkZjA3NS1hZjI2LTRjZTQtOTU0YS03OGU3MDBkOWYyYjQiLCJ0b2tlbl91c2UiOiJpZCIsImF1dGhfdGltZSI6MTY3NTA5MTg2NywibmFtZSI6InBhdWwiLCJleHAiOjE2NzUxNzgyNjcsImlhdCI6MTY3NTA5MTg2NywianRpIjoiM2JiZWJkNzctMGY4Zi00NDE3LTgzMmYtNDlkZjBiZDU5OWFhIiwiZW1haWwiOiJhbmthbGFnb0BnbWFpbC5jb20ifQ.JJzZUInlsh5TGM-X5En-V5EaNa4bLlCrBCDh1jPDLAGoIFZE_Crv169RJ8vO032Ran_DURroMrIpB1_GIp3_HanPLD-Zh3IC8XCx_62jwY4_dcTqs003Jr6kwRRI35ydtOGf10xQ3DBLjTXvA4oKAn8IlrsbustKrBK1WzeCFSFezzll0Iu_KYXhnlaLA-bSAq3bGJPkubzL1BUiJY0ONDZCyXaHG6G3T3znh1q80u9xMIYpE5gENNGSA0yA6LwLq1ZM_q9rZN15q8BfwUSKSo4dXYCaqtxEPWr5TAgx9l3PPwBvjyN_4XNxgG9oOAa4SYeXgNtT8wGGEdMjt1fvhQ")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("err: ", err)
	}
	s := string(body)
	var feedRequest entities.Request
	errEncode := json.Unmarshal([]byte(s), &feedRequest)
	if errEncode != nil {
		log.Fatal("errEncode: ", errEncode)
	}
	return feedRequest
}
