package entities

type Error struct {
	StatusCode  int    `json:"status_code"`
	ErrorCode   int    `json:"error_code"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
