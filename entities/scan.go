package entities

type cacheable interface {
	Folder | Item
}

type Folder struct {
	Name  string `json:"name"`
	Items []any  `json:"items"`
}

type File struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type Node[T cacheable] struct {
	Name string `json:"name"`

	Items []T `json:"items"`
}
