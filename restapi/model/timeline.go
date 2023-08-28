package model

type Timeline struct {
	Posts      []*Post     `json:"posts"`
	Pagenation *Pagenation `json:"pagination"`
}

type Pagenation struct {
	NextCursor     string `json:"nextCursor"`     // 次のページのカーソル
	PreviousCursor string `json:"previousCursor"` // 前のページのカーソル
	Limit          int    `json:"limit"`          // 1ページあたりの件数
}
