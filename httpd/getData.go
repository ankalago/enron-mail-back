package httpd

import (
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
				"sql": "SELECT * FROM \"enron_mail\"%s",
				"start_time": 1672549200000000,
				"end_time": 1675227540000000,
				"from": 0,
				"size": 1000
			}
		}`, where)
	req, err := http.NewRequest("POST", "https://api.zincsearch.com/api/enterprise_16O476E9609y7Lz/_search", strings.NewReader(query))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", "Bearer eyJraWQiOiJ1TVJxZlRhUmp1c1pyWUZFRWVFMlVFeUVhQW50WVhMbEZOQU11eitIZkZFPSIsImFsZyI6IlJTMjU2In0.eyJhdF9oYXNoIjoiNlQxQWYyS1I4a0lKQi15TFNsaUFOZyIsInN1YiI6ImU0MjY1YjM4LWYyODctNDI0Mi05YTIzLTAzZjE4NzlhMGM2MCIsImNvZ25pdG86Z3JvdXBzIjpbInVzLWVhc3QtMl9weTJnSVlZOXZfR29vZ2xlIl0sImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwiaXNzIjoiaHR0cHM6XC9cL2NvZ25pdG8taWRwLnVzLWVhc3QtMi5hbWF6b25hd3MuY29tXC91cy1lYXN0LTJfcHkyZ0lZWTl2IiwiY29nbml0bzp1c2VybmFtZSI6Imdvb2dsZV8xMDc1MTQxNjU4MTYwMzQ2MzAxODciLCJnaXZlbl9uYW1lIjoiUGHDumwiLCJub25jZSI6ImFVVVNYQnc3UGEwR2l6Zy0yaE5KWlhsUzE1VFM0ZnBoQmpMWV9UVXE1X3NiRGx3ZFVZVm0yajRYR0tJX2piVElIUTVvcUpyYVVfXzNTTGFlV1pQY1Rrc3EwajN2Ty1lWDhtaFF2M05jSGxQTURNemR4UkZvOFdKanV5RjJfT3l3Qm9CeDZYVE9lTF9nWGY1NEN0WTZqR1M5VzZRcm5fLVpLWk5UY1UwaFFUSSIsInBpY3R1cmUiOiJodHRwczpcL1wvbGgzLmdvb2dsZXVzZXJjb250ZW50LmNvbVwvYVwvQUVkRlRwNUNoenBoNTlfeG11Q184OEJCTUxZb001ckEwcVYtaERUVnhZb083RkU9czk2LWMiLCJvcmlnaW5fanRpIjoiZGEwZGFjOGYtMDRjOS00MDM5LTlhNjgtMzIzNjA3MjQ2OGMyIiwiYXVkIjoiM2FzZGNkdHRpZ25wZjVxanM3bm10NGY2NjYiLCJpZGVudGl0aWVzIjpbeyJ1c2VySWQiOiIxMDc1MTQxNjU4MTYwMzQ2MzAxODciLCJwcm92aWRlck5hbWUiOiJHb29nbGUiLCJwcm92aWRlclR5cGUiOiJHb29nbGUiLCJpc3N1ZXIiOm51bGwsInByaW1hcnkiOiJ0cnVlIiwiZGF0ZUNyZWF0ZWQiOiIxNjc0MDg4MTEzNzQ2In1dLCJ0b2tlbl91c2UiOiJpZCIsImF1dGhfdGltZSI6MTY3NDc5MzY0NCwibmFtZSI6IlBhw7psIErDoWNvbWUiLCJleHAiOjE2NzQ4ODAwNDQsImlhdCI6MTY3NDc5MzY0NCwiZmFtaWx5X25hbWUiOiJKw6Fjb21lIiwianRpIjoiNWUwYWQ2ZGYtNzY1Ny00MDVmLTllNjctYTZkOWM5MWRjNTZlIiwiZW1haWwiOiJhbmthbGFnb0BnbWFpbC5jb20ifQ.Yk8TaLP7H8OdynM2b8rzYRyDABkdLC4aguLCVXWp4P7pIhB7zY49rlWt4whE6coUyCVr5CDhSIHWMeF7Ak_rfjrd1sHzYXhGHkeP-95OSezUN7HRLXl2OvQwOwydQUeSZzmcaRoxsBXkyU1_8pNJ35aX4JOkal43FWSTFuNbPA4gSZ0fexkj36qG1J_djbsw7z2JYEXJ0svXTZixUhi0oTegOFxjR2dwXxqhPIRPUnjFnfLtEG35GB6RZaCs9PRvb5NHgDHh4ZWoa9_cZVqLYo_-6DbAIUcx5UeBRXlzViCMIuwq1NDoYi0OvgLXBec0cD76_6eLLUfs4DZCJ2Tk1g")
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
