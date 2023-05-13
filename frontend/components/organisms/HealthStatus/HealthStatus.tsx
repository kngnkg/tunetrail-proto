import React from "react"
import { Health } from "@/types"

import HealthIndicator from "@/components/atoms/HealthIndicator/HealthIndicator"

interface HealthStatusProps {
  health: Health
}

// ヘルスチェックの結果を表示するコンポーネント
export default function HealthStatus({ health }: HealthStatusProps) {
  return (
    <div className="space-y-4">
      <div>
        <h3 className="text-xl">Health Check:</h3>
        <HealthIndicator status={health.health} />
      </div>
      <div>
        <h3 className="text-xl">Database Check:</h3>
        <HealthIndicator status={health.database} />
      </div>
    </div>
  )
}
