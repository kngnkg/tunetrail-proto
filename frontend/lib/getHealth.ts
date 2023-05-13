import { ApiContext, Health } from "@/types"

import { fetchJson } from "./fetchJson"
import { HealthSchema } from "./validations/health.schema"

// getHealthはAPIの稼働状態を取得する
export const getHealth = async (context: ApiContext): Promise<Health> => {
  const data: unknown = await fetchJson(`${context.apiRoot}/health`)
  if (!data) {
    throw new Error("Failed to fetch data from the API.")
  }

  const validationResult = HealthSchema.safeParse(data)
  if (!validationResult.success) {
    throw new Error(validationResult.error.message)
  }

  return validationResult.data
}
