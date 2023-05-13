/**
 * 指定されたリソースからデータを取得し、JSONレスポンスを返す。
 *
 * @param resource - 取得するリソースのURL
 * @param init - フェッチリクエストをカスタマイズするためのオプションのRequestInitオブジェクト
 * @returns APIからのJSONデータを解決するPromise、またはリクエストが失敗した場合はエラーをスローする
 * @throws リクエストが失敗した場合、またはレスポンスがJSONとして解析できない場合にエラーをスローする
 */
export const fetchJson = async (
  resource: RequestInfo,
  init?: RequestInit
): Promise<unknown> => {
  const response = await fetch(resource, init)
  if (!response.ok) {
    const errorResponse = await response.json()
    const error = new Error(
      errorResponse.message ?? "Failed to fetch data from the API."
    )
    throw error
  }

  return response.json()
}
