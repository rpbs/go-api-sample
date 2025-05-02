package main

type Book struct {
	BookID      int     `json:"book_id"`
	Title       string  `json:"title"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	TotalCopies int     `json:"total_copies"`
	CopiesInUse int     `json:"copies_in_use"`
	Type        *string `json:"type"`     // Nullable field
	ISBN        *string `json:"isbn"`     // Nullable field
	Category    *string `json:"category"` // Nullable field
}
