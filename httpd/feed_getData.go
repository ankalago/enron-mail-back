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

func NewsfeedGet(queryText string) newsfeed.Request {
	return getData(queryText)
}

func getData(queryText string) newsfeed.Request {
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
	req.Header.Set("authorization", "Bearer eyJraWQiOiJ1TVJxZlRhUmp1c1pyWUZFRWVFMlVFeUVhQW50WVhMbEZOQU11eitIZkZFPSIsImFsZyI6IlJTMjU2In0.eyJhdF9oYXNoIjoiaGsyWFNDRkhhRnJBUkUxVTZMTDFudyIsInN1YiI6ImU0MjY1YjM4LWYyODctNDI0Mi05YTIzLTAzZjE4NzlhMGM2MCIsImNvZ25pdG86Z3JvdXBzIjpbInVzLWVhc3QtMl9weTJnSVlZOXZfR29vZ2xlIl0sImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwiaXNzIjoiaHR0cHM6XC9cL2NvZ25pdG8taWRwLnVzLWVhc3QtMi5hbWF6b25hd3MuY29tXC91cy1lYXN0LTJfcHkyZ0lZWTl2IiwiY29nbml0bzp1c2VybmFtZSI6Imdvb2dsZV8xMDc1MTQxNjU4MTYwMzQ2MzAxODciLCJnaXZlbl9uYW1lIjoiUGHDumwiLCJwaWN0dXJlIjoiaHR0cHM6XC9cL2xoMy5nb29nbGV1c2VyY29udGVudC5jb21cL2FcL0FFZEZUcDVDaHpwaDU5X3htdUNfODhCQk1MWW9NNXJBMHFWLWhEVFZ4WW9PN0ZFPXM5Ni1jIiwib3JpZ2luX2p0aSI6ImNjNjRmN2FlLThjZTctNDc2Zi04ODhkLTVmYmQ1Yzc0OGNiMyIsImF1ZCI6IjNhc2RjZHR0aWducGY1cWpzN25tdDRmNjY2IiwiaWRlbnRpdGllcyI6W3sidXNlcklkIjoiMTA3NTE0MTY1ODE2MDM0NjMwMTg3IiwicHJvdmlkZXJOYW1lIjoiR29vZ2xlIiwicHJvdmlkZXJUeXBlIjoiR29vZ2xlIiwiaXNzdWVyIjpudWxsLCJwcmltYXJ5IjoidHJ1ZSIsImRhdGVDcmVhdGVkIjoiMTY3NDA4ODExMzc0NiJ9XSwidG9rZW5fdXNlIjoiaWQiLCJhdXRoX3RpbWUiOjE2NzQwODgxMTQsIm5hbWUiOiJQYcO6bCBKw6Fjb21lIiwiZXhwIjoxNjc0NjkwMzUzLCJpYXQiOjE2NzQ2MDM5NTMsImZhbWlseV9uYW1lIjoiSsOhY29tZSIsImp0aSI6ImM4OGE0MWI5LWY1MGMtNGE2Yi1iMmM0LTkwYTNjNWJkODRlNCIsImVtYWlsIjoiYW5rYWxhZ29AZ21haWwuY29tIn0.jXorMMOe7MlQCCRMnkguWnnWLvxR8weKtu6XNmXLVMmlTX4-BSb4Hj52b1dnM-qwHeFFTjPNkYbvw9y6T_8rimDEE5mS_U_YSWAhUl0OrxIQKoqMLyBIEspOqRKWyzfoSnuuvh9K2Mr8xGMPCEOfQHT5VFUPHXg_VnE-6eYlvWh8v6l_Crvk3ObuFX4HHEW_TUNxFgsEfEr3mBYV7YCGktaJ6u_EtqaO6GELQiebdS9gecQh_W12ckKhNx_LeSqcnzYinir7dHILXpWMIXuTKYKk2x3n81kM8UHjFlOo25igAHelP3CHja0LGWAwkMtWF67j3LOfpQCxpsm_Xo8_WA")
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
