export type FetchError = string

// APIのエラー時のレスポンスの型
export type ApiError = {
  code: number
  developerMessage: string
  userMessage: string
}

export const isApiError = (arg: any): arg is ApiError => {
  return (
    arg.code !== undefined &&
    arg.developerMessage !== undefined &&
    arg.userMessage !== undefined
  )
}
