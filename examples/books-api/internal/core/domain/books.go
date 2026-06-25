package domain

type GetBook struct {
	Bookid string `params:"bookId"`
}
type BorrowBook struct {
	ID       string `json:"id"`
	Bookid   string `params:"bookId"`
	MemberID string `json:"member_id"`
}
type ReturnBook struct {
	ID     string `json:"id"`
	Bookid string `params:"bookId"`
}
type HealthCheck struct {
}
type ListBooks struct {
	Status string `query:"status"`
	Limit  int    `query:"limit"`
	Offset int    `query:"offset"`
}
type CreateBook struct {
	ID     string `json:"id"`
	Author string `json:"author"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
}
