package types

type Image struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Format string `json:"format"`
}

type Video struct {
	URL      string `json:"url"`
	Format   string `json:"format"`
	Duration int    `json:"duration"`
	Size     int    `json:"size"`
	Image    *Image `json:"image"`
}

type Audio struct {
	URL      string `json:"url"`
	Format   string `json:"format"`
	Duration int    `json:"duration"`
	Size     int    `json:"size"`
}

type File struct {
	URL    string `json:"url"`
	Name   string `json:"name"`
	Size   int    `json:"size"`
	Format string `json:"format"`
}

type WebPage struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Image   *Image `json:"image"`
	URL     string `json:"url"`
}
