import { FetchError, isApiError } from "@/types/error"
import { MESSAGE } from "@/config/messages"

export interface ConfirmSignupData {
  userName: string
  code: string
}
export const isConfirmSignupData = (arg: any): arg is ConfirmSignupData => {
  return arg.userName !== undefined && arg.code !== undefined
}

export const confirmSignup = async (
  apiRoot: string,
  data: ConfirmSignupData
): Promise<null | FetchError> => {
  try {
    const response = await fetch(`${apiRoot}/auth/confirm`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    })

    if (!response.ok) {
      const errorResponse = await response.json()
      if (!isApiError(errorResponse)) {
        const error = new Error(
          errorResponse.message ??
            `Failed to fetch data from the API. Status: ${response.status}`
        )
        console.error(error)
        throw error
      }

      return errorResponse.userMessage
    }
    return null
  } catch (e) {
    console.error(e)
    if (e instanceof Error) {
      return e.message
    }

    return MESSAGE.UNKNOWN_ERROR
  }
}
