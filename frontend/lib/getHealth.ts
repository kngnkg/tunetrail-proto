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
      console.error(response)
      const errorResponse = await response.json()
      const error = new Error(
        errorResponse.message ??
          `Failed to fetch data from the API. Status: ${response.status}`
      )
      throw error
    }

    const data = await response.json()
    if (!data) {
      throw new Error("data is empty from the API.")
    }

    const validationResult = HealthSchema.safeParse(data)
    if (!validationResult.success) {
      throw new Error(validationResult.error.message)
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
