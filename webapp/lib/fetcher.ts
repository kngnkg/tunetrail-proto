import { refreshToken } from "@/services/auth/refreshToken"

import { env } from "@/env.mjs"
import ApiError, { isApiError } from "@/types/error"

const tokenExpiredCode = 4105

export const fetcher = async (
  resource: RequestInfo,
  init?: RequestInit
): Promise<any> => {
  try {
    const response = await fetch(resource, init)

    if (!response.ok) {
      const errResp = await response.json()

      if (!isApiError(errResp)) {
        // エラーレスポンスがAPIエラーでない場合はエラーを投げる
        throw new Error(
          errResp.message ??
            `Failed to fetch data from the API. Status: ${response.status}`
        )
      }

      // 認証エラーの場合
      if (response.status === 401) {
        switch (errResp.code) {
          case tokenExpiredCode:
            // トークンが期限切れの場合はリフレッシュトークンを送信して再度リクエストを送る
            await refreshToken(env.NEXT_PUBLIC_AUTH_API_ROOT)
            return fetcher(resource, init)
          // case 4106:
          // リフレッシュトークンが期限切れの場合はログイン画面に遷移する
          // case 4107:
          // CSRFトークンが期限切れの場合は...
        }
      }

      throw new ApiError(
        errResp.code,
        errResp.developerMessage,
        errResp.userMessage
      )
    }

    return response.json()
  } catch (e) {
    throw e
  }
}

export default fetcher
