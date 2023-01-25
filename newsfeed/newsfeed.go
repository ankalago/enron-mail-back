package newsfeed

type Request struct {
	Took     int    `json:"took"`
	Hits     []Item `json:"hits"`
	Total    int    `json:"total"`
	From     int    `json:"from"`
	Size     int    `json:"size"`
	ScanSize int    `json:"scan_size"`
}

type Item struct {
	Athlete    string `json:"athlete"`
	City       string `json:"city"`
	Country    string `json:"country"`
	Discipline string `json:"discipline"`
	Event      string `json:"event"`
	Gender     string `json:"gender"`
	Medal      string `json:"medal"`
	Season     string `json:"season"`
	Sport      string `json:"sport"`
	Year       int    `json:"year"`
	Timestamp  int    `json:"_timestamp"`
}

type Repo struct {
	Items []Item
}

func New() *Repo {
	return &Repo{
		Items: []Item{},
	}
}

func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

func (r *Repo) AddAll(items []Item) {
	r.Items = items
}

func (r *Repo) GetAll() []Item {
	return r.Items
}
