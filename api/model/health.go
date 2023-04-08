package model

const (
	// 正常に稼働中
	StatusGreen = "green"
	// 一部正常に稼働していない
	StatusOrange = "orange"
	// 正常に稼働していない
	StatusRed = "red"
)

// Healthはサービスの稼働状態を表す構造体
type Health struct {
	Health   string `json:"health"`
	Database string `json:"database"`
}
