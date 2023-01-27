package entities

type Request struct {
	Took     int    `json:"took"`
	Hits     []Item `json:"hits"`
	Total    int    `json:"total"`
	From     int    `json:"from"`
	Size     int    `json:"size"`
	ScanSize int    `json:"scan_size"`
}

type Item struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Subject   string `json:"subject"`
	Origin    string `json:"origin"`
	Content   string `json:"content"`
	Timestamp int    `json:"_timestamp"`
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
