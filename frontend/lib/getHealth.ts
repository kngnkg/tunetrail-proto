import { ApiContext, Health } from "@/types"

import { HealthSchema } from "./validations/health.schema"

// getHealthはAPIの稼働状態を取得する
export const getHealth = async (context: ApiContext): Promise<Health> => {
  console.log(`fetching ${context.apiRoot}/health`)
  try {
    const response = await fetch(`${context.apiRoot}/health`, {
      cache: "no-store",
    })
    if (!response.ok) {
      const errorResponse = await response.json()
      const error = new Error(
        errorResponse.message ??
          `Failed to fetch data from the API. Status: ${response.status}`
      )
      console.error(error)
      throw error
    }

    const data = await response.json()
    if (!data) {
      const error = new Error("data is empty from the API.")
      console.error(error)
      throw error
    }

    const validationResult = HealthSchema.safeParse(data)
    if (!validationResult.success) {
      const error = new Error(validationResult.error.message)
      console.error(error)
      throw error
    }

    return validationResult.data
  } catch (e) {
    if (e instanceof Error) {
      console.error(e.message)
    } else {
      console.error(e)
    }

    return {
      health: "red",
      database: "red",
    }
  }
}
