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
	req.Header.Set("authorization", "Bearer eyJraWQiOiJ1TVJxZlRhUmp1c1pyWUZFRWVFMlVFeUVhQW50WVhMbEZOQU11eitIZkZFPSIsImFsZyI6IlJTMjU2In0.eyJhdF9oYXNoIjoic212NXlZOE1QcmhOZnJyc1dURU5VQSIsInN1YiI6ImQxM2NjMDA1LWNmZGEtNDM0Ny05YWU4LWQxMzIxZmM1ODAzZCIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJpc3MiOiJodHRwczpcL1wvY29nbml0by1pZHAudXMtZWFzdC0yLmFtYXpvbmF3cy5jb21cL3VzLWVhc3QtMl9weTJnSVlZOXYiLCJjb2duaXRvOnVzZXJuYW1lIjoiZDEzY2MwMDUtY2ZkYS00MzQ3LTlhZTgtZDEzMjFmYzU4MDNkIiwib3JpZ2luX2p0aSI6IjQ4OTU1YzBiLTUwZTktNDRlMy05YmJiLWEzZDExYjhhZGRkMCIsImF1ZCI6IjNhc2RjZHR0aWducGY1cWpzN25tdDRmNjY2IiwiZXZlbnRfaWQiOiI3NzJhZjgyYi0zODI4LTRlYjktYjcwNy1jMGJmZTQ4ZWE5MTIiLCJ0b2tlbl91c2UiOiJpZCIsImF1dGhfdGltZSI6MTY3NTE4NjQ1NCwibmFtZSI6InBhdWwiLCJleHAiOjE2NzUyNzI4NTQsImlhdCI6MTY3NTE4NjQ1NCwianRpIjoiY2E3Y2YxZjUtMGVhOC00NzEzLWEzNzQtNGMzYmJjMzVhNmFmIiwiZW1haWwiOiJhbmthbGFnb0BnbWFpbC5jb20ifQ.iShiAkSflvOoa8BUOIeBkuaPuRukF-n8_CSNO6nQQIpp-mFsVTvoEV7Nmf0_fC-LSB4WZh9wBB2klTLM7ulWL8AiJj6Fr81hAAcUt9HkfNHFVSoU2pO7n5Al3YOG0ozu7Yy54npyXPJjcoPGUq-W-a7C_aVQGc2mC6EbhXM6SIxAOv_TuacDpIheVq96Q29xT45wiwI_wxyNGLm7o7i05gsQK7GLUeksMXzxu6W9DXfU6Hc3kk2d-Of-yOlZU55P7AiP-649KizdlWEfOIOrM2eRy_H8WTR0f8FimOO1T_xR_kdO2gBN5_MzbFI_JO3u-fMfbnWWebr5hSEcexVBlw")
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
