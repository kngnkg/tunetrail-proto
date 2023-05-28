import { ApiContext, Health } from "@/types"

import { HealthSchema } from "./validations/health.schema"

// getHealthはAPIの稼働状態を取得する
export const getHealth = async (context: ApiContext): Promise<Health> => {
  console.log(`fetching ${context.serverApiRoot}/health`)
  const response = await fetch(`${context.serverApiRoot}/health`, {
    cache: "no-store",
  })
  if (!response.ok) {
    console.error(response)
    const errorResponse = await response.json()
    const error = new Error(
      errorResponse.message ?? "Failed to fetch data from the API."
    )
    throw error
  }

  const data = await response.json()
  if (!data) {
    console.error(`empty response from ${context.serverApiRoot}/health`)
    console.error(response)
    throw new Error("Failed to fetch data from the API.")
  }

  const validationResult = HealthSchema.safeParse(data)
  if (!validationResult.success) {
    throw new Error(validationResult.error.message)
  }

  return validationResult.data
}
