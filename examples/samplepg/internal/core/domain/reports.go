package domain

type GetReports struct {
	Status string `query:"status"`
	Limit  int    `query:"limit"`
	Offset int    `query:"offset"`
}
type CreateReport struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Title       string `json:"title"`
}
