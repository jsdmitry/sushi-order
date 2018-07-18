package model

// MenuItem contains Caption, ImageUrl and Description
type MenuItem struct {
	Caption     string
	ImageURL    string
	Description string
	Price       uint64
}

// CategoryItem contains Caption, ImageUrl and MenuUrl
type CategoryItem struct {
	ID       int
	Caption  string
	ImageURL string
	MenuURL  string
}
