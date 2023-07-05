import { ApiContext } from "@/types/api-context"
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
 * @param context // APIのルートパス
 * @param data // サインアップに必要なデータ
 * @returns // エラーがあればエラーメッセージを返す
 */
export const signup = async (
  context: ApiContext,
  data: SignupData
): Promise<null | FetchError> => {
  try {
    const response = await fetch(`${context.apiRoot}/user/register`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    })

    if (!response.ok) {
      const errorResponse = await response.json()
      console.error(errorResponse)
      if (!isApiError(errorResponse)) {
        const error = new Error(
          errorResponse.message ??
            `Failed to fetch data from the API. Status: ${response.status}`
        )
        console.error(error)
        throw error
      }

      switch (response.status) {
        case 400:
          let errorMsg = MESSAGE.UNKNOWN_ERROR
          if (errorResponse.msg === "ユーザー名が既に登録されています。") {
            errorMsg = MESSAGE.DUP_USERNAME
          }
          if (errorResponse.msg === "メールアドレスが既に登録されています。") {
            errorMsg = MESSAGE.DUP_EMAIL
          }

          return errorMsg
        case 500:
          return MESSAGE.UNKNOWN_ERROR
        default:
          return MESSAGE.UNKNOWN_ERROR
      }
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
