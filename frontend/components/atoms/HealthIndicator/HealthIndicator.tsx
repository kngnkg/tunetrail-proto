interface HealthIndicatorProps {
  status: "green" | "orange" | "red"
}

// ヘルスチェックの結果を表示するコンポーネント
export default function HealthIndicator({ status }: HealthIndicatorProps) {
  // statusが green の場合のみ、テキストの色を緑にする
  const statusColor = status === "green" ? "text-green-500" : "text-red-500"

  return (
    <div className={`font-bold text-2xl ${statusColor}`}>
      {status.toUpperCase()}
    </div>
  )
}
