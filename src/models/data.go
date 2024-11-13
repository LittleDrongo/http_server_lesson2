package models

import "slices"

var DB []Book

type Book struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Author        Author `json:"author"`
	YearPublished int    `json:"year_published"`
}

type Author struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	BornYear int    `json:"born_year"`
}

func init() {

	DB = append(DB, Book{
		ID:            1,
		Title:         "Lord of the Rings. Vol. 1",
		YearPublished: 1978,
		Author: Author{
			Name:     "J.R.",
			LastName: "Tolkin",
			BornYear: 1892,
		}},
		Book{
			ID:            2,
			Title:         "The Two Towers",
			YearPublished: 1978,
			Author: Author{
				Name:     "J.R.",
				LastName: "Tolkin",
				BornYear: 1954,
			}},
	)

}

// FindBookByID function find book by id.
func FindBookByID(id int) (book Book, found bool) {
	for _, b := range DB {
		if b.ID == id {
			book = b
			found = true
			break
		}
	}

	return book, found
}

func DeleteBookById(id int) (found bool) {
	for idx, b := range DB {
		if b.ID == id {
			DB = slices.Delete(DB, idx, idx+1)
			found = true
			return found
		}

	}
	return found
}

func UpdateBookById(id int, book Book) (found bool) {
	for k, b := range DB {
		if b.ID == id {
			if len(b.Title) > 0 {
				DB[k].Title = b.Title
			}
			if len(b.Author.Name) > 0 {
				DB[k].Author.Name = book.Author.Name
			}

			if len(b.Author.LastName) > 0 {
				DB[k].Author.LastName = book.Author.LastName
			}

			if b.Author.BornYear > 0 {
				DB[k].Author.BornYear = book.Author.BornYear
			}

			if b.YearPublished > 0 {
				DB[k].YearPublished = book.YearPublished
			}
			found = true
		}

	}
	return found
}
