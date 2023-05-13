import React from "react"

import { apiContext } from "@/lib/apiContext"
import { getHealth } from "@/lib/getHealth"
import HealthStatus from "@/components/organisms/HealthStatus/HealthStatus"

// ヘルスチェックの結果を表示するページ
export default async function Health() {
  const health = await getHealth(apiContext)

  return (
    <div className="container mx-auto p-8">
      <h1 className="text-3xl mb-8">Health Check Page</h1>
      <HealthStatus health={health} />
    </div>
  )
}
