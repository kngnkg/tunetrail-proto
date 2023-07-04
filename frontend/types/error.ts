export type FetchError = string

// APIのエラー時のレスポンスの型
export type ApiError = {
  msg: string
}

export const isApiError = (arg: any): arg is ApiError => {
  return arg.msg !== undefined
}
