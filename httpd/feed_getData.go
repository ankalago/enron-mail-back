package httpd

import (
	"chi-example/newsfeed"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func GetData(queryText string) newsfeed.Request {
	query := fmt.Sprintf(`{
			"query": {
				"sql": "SELECT * FROM 'olympics' WHERE Event LIKE '%s'",
				"start_time": 1674517800000000,
				"end_time": 1674517920000000,
				"from": 0,
				"size": 1000
			},
			"aggs": {
				"histogram": "select histogram(\"_timestamp\", '1 second') AS key, count(*) AS num from query GROUP BY key ORDER BY key"
			}
		}`, queryText+"%")
	req, err := http.NewRequest("POST", "https://api.zincsearch.com/api/enterprise_16O476E9609y7Lz/_search", strings.NewReader(query))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", "Bearer eyJraWQiOiJ1TVJxZlRhUmp1c1pyWUZFRWVFMlVFeUVhQW50WVhMbEZOQU11eitIZkZFPSIsImFsZyI6IlJTMjU2In0.eyJhdF9oYXNoIjoiOHdKV0J4UXBJZFhpb1NPeWNqOURrZyIsInN1YiI6ImU0MjY1YjM4LWYyODctNDI0Mi05YTIzLTAzZjE4NzlhMGM2MCIsImNvZ25pdG86Z3JvdXBzIjpbInVzLWVhc3QtMl9weTJnSVlZOXZfR29vZ2xlIl0sImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwiaXNzIjoiaHR0cHM6XC9cL2NvZ25pdG8taWRwLnVzLWVhc3QtMi5hbWF6b25hd3MuY29tXC91cy1lYXN0LTJfcHkyZ0lZWTl2IiwiY29nbml0bzp1c2VybmFtZSI6Imdvb2dsZV8xMDc1MTQxNjU4MTYwMzQ2MzAxODciLCJnaXZlbl9uYW1lIjoiUGHDumwiLCJub25jZSI6InpWUVFYTkctajU5SmZsSXRXUlpreEE0SkxFMGRrcVlvWFVFcEJxN29WVjlYTjMyV1VHSHRRbEFjZTFwaWstUVdfRnd2YnZkMDVqQ19ZOVlMWDZ6dWZKN2NKVWF3OU5ldFdOeHV2aVBqYU95S3kwU2JlelNiMWhiSU8tMUltOTdoVFI1bFFSc25idmM3X0dZYS0tLTE0NkVnM1FQMlZFS01YY2lXbVZ5Sk1hOCIsInBpY3R1cmUiOiJodHRwczpcL1wvbGgzLmdvb2dsZXVzZXJjb250ZW50LmNvbVwvYVwvQUVkRlRwNUNoenBoNTlfeG11Q184OEJCTUxZb001ckEwcVYtaERUVnhZb083RkU9czk2LWMiLCJvcmlnaW5fanRpIjoiNDllN2Y5ODUtMTdlOS00ZWE4LTg5ZTctY2RlNDkzZDk3Y2JjIiwiYXVkIjoiM2FzZGNkdHRpZ25wZjVxanM3bm10NGY2NjYiLCJpZGVudGl0aWVzIjpbeyJ1c2VySWQiOiIxMDc1MTQxNjU4MTYwMzQ2MzAxODciLCJwcm92aWRlck5hbWUiOiJHb29nbGUiLCJwcm92aWRlclR5cGUiOiJHb29nbGUiLCJpc3N1ZXIiOm51bGwsInByaW1hcnkiOiJ0cnVlIiwiZGF0ZUNyZWF0ZWQiOiIxNjc0MDg4MTEzNzQ2In1dLCJ0b2tlbl91c2UiOiJpZCIsImF1dGhfdGltZSI6MTY3NDcwMTc0MywibmFtZSI6IlBhw7psIErDoWNvbWUiLCJleHAiOjE2NzQ3ODgxNDMsImlhdCI6MTY3NDcwMTc0MywiZmFtaWx5X25hbWUiOiJKw6Fjb21lIiwianRpIjoiMDZiZDJmN2ItODYyNC00ZjllLTgxNTItYmNkZDU5NDNhNmUwIiwiZW1haWwiOiJhbmthbGFnb0BnbWFpbC5jb20ifQ.CgITBQkcdg13wsq2o0NCgUfgqgT_RJeUZO8sF_HQuthm7qV_PYESBtVrsxR6VcQOxDlibu4fFNPrGiYu8SAA2YXy0V69JhVmltzs1SvsovKH-v_YEjv6PMmVnvU_CK1TvvXyqJuv_VZk4rcmsHRKYyswxgRTks1uQ2E58I8nbFRZWWgjNxstwLbzWqMEtst3CHu0SDUU_50HyxE7hiteTdV-DW4wKzoJHNrSdpnTu9ZjRbbuU-OhCspPJVx1wFgS3gowipf6lDI8WKnPHb8ACR7vC5w__EpvBNP6574M1CNqtNvK6HEJwohmuij72DTF-rnEmHYx_G8KBlnK5M9r9Q")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	s := string(body)
	var myStoredVariable newsfeed.Request
	errEncode := json.Unmarshal([]byte(s), &myStoredVariable)
	if errEncode != nil {
		log.Fatal(errEncode)
	}
	return myStoredVariable
}
