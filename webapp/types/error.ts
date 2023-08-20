export type FetchError = string

// // APIのエラー時のレスポンスの型
// export type ApiError extends Error = {
//   code: number
//   developerMessage: string
//   userMessage: string
// }

export default class ApiError extends Error {
  code: number
  developerMessage: string
  userMessage: string

  constructor(code: number, developerMessage: string, userMessage: string) {
    super(developerMessage)
    this.code = code
    this.developerMessage = developerMessage
    this.userMessage = userMessage
  }
}

export const isApiError = (arg: any): arg is ApiError => {
  return (
    arg.code !== undefined &&
    arg.developerMessage !== undefined &&
    arg.userMessage !== undefined
  )
}
