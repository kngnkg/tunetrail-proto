import { isApiError } from "@/types/error"

export const fetcher = async (
  resource: RequestInfo,
  init?: RequestInit
): Promise<any> => {
  try {
    const response = await fetch(resource, init)

    if (!response.ok) {
      const errResp = await response.json()

      if (isApiError(errResp)) {
        // TODO: エラーコードによる分岐等
        throw new Error(errResp.userMessage)
      }

      throw new Error(
        errResp.message ??
          `Failed to fetch data from the API. Status: ${response.status}`
      )
    }

    return response.json()
  } catch (e) {
    throw e
  }
}

export default fetcher
