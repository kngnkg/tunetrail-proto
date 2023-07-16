import dns from "dns"

import { Health } from "@/types/health"
import { HealthSchema } from "@/lib/validations/health.schema"

// getHealthはAPIの稼働状態を取得する
export const getHealth = async (apiRoot: string): Promise<Health> => {
  // DNSのlookupを行う
  const apiHost = apiRoot.replace(/^https:\/\//, "")
  console.log(`looking up ${apiHost}`)
  dns.lookup(apiHost, (err, address, family) => {
    if (err) {
      console.error(err)
      return
    }
    console.log(`address: ${address} family: IPv${family}`)
  })

  console.log(`fetching ${apiRoot}/health`)
  try {
    const response = await fetch(`${apiRoot}/health`, {
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
      console.error(e)
      // console.error(e.message)
    } else {
      console.error(e)
    }

    return {
      health: "red",
      database: "red",
    }
  }
}
