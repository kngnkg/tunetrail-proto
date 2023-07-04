// APIの稼働状態を表す
export type Health = {
  health: "green" | "orange" | "red"
  database: "green" | "orange" | "red"
}
