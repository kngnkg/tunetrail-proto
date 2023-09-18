package model

type Timeline struct {
	Posts      []*Post     `json:"posts"`
	Pagination *Pagination `json:"pagination"`
}

type Pagination struct {
	NextCursor     string `json:"nextCursor"`     // 次のページのカーソル
	PreviousCursor string `json:"previousCursor"` // 前のページのカーソル
	Limit          int    `json:"limit"`          // 1ページあたりの件数
}
