package main

import "chi-example/httpd"

func main() {
	data := `{ "index" : { "_index" : "olympics" } }
	{"Year": 1896, "City": "Quito", "Sport": "Aquatics", "Discipline": "Swimming", "Athlete": "HAJOS, Alfred", "Country": "HUN", "Gender": "Men", "Event": "100M Freestyle", "Medal": "Gold", "Season": "summer"}
	{ "index" : { "_index" : "olympics" } }
	{"Year": 1896, "City": "Quito", "Sport": "Aquatics", "Discipline": "Swimming", "Athlete": "HERSCHMANN, Otto", "Country": "AUT", "Gender": "Men", "Event": "100M Freestyle", "Medal": "Silver", "Season": "summer"}
	{ "index" : { "_index" : "olympics" } }
	{"Year": 1896, "City": "Quito", "Sport": "Aquatics", "Discipline": "Swimming", "Athlete": "DRIVAS, Dimitrios", "Country": "GRE", "Gender": "Men", "Event": "100M Freestyle For Sailors", "Medal": "Bronze", "Season": "summer"}` // string | Query

	httpd.BulkData(data, 10)
}
