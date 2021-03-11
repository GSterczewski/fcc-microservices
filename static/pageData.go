package static

// Link - html <a> data
type Link struct {
	Name string
	Href string
}

//PageData defines html data to render
type PageData struct {
	Title string
	Links []Link
}
