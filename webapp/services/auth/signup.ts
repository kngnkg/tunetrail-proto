import { FetchError, isApiError } from "@/types/error"
import { MESSAGE } from "@/config/messages"

export interface SignupData {
  userName: string
  name: string
  password: string
  email: string
}
export const isSignupData = (arg: any): arg is SignupData => {
  return (
    arg.userName !== undefined &&
    arg.name !== undefined &&
    arg.password !== undefined &&
    arg.email !== undefined
  )
}

/**
 * サインアップ処理
 * @param apiRoot // APIのルートパス
 * @param data // サインアップに必要なデータ
 * @returns // エラーがあればエラーメッセージを返す
 */
export const signup = async (
  apiRoot: string,
  data: SignupData
): Promise<null | FetchError> => {
  try {
    const response = await fetch(`${apiRoot}/auth/register`, {
      method: "POST",
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

      // switch (response.status) {
      //   case 400:
      //     return errorResponse.userMessage
      //   case 409:
      //     return errorResponse.userMessage
      //   case 500:
      //     return errorResponse.userMessage
      //   default:
      //     return MESSAGE.UNKNOWN_ERROR
      // }

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
