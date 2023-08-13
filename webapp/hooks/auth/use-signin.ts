import * as React from "react"

import { isApiError } from "@/types/error"
import { isUser } from "@/types/user"
import { MESSAGE } from "@/config/messages"

import { useSignedInUser } from "./use-signedin-user"

export interface SigninData {
  userName: string
  email: string
  password: string
}
export const isSigninData = (arg: any): arg is SigninData => {
  return (
    (arg.userName !== undefined || arg.email !== undefined) &&
    arg.password !== undefined
  )
}

export const useSignin = () => {
  const { setSignedInUser } = useSignedInUser()
  const [error, setError] = React.useState<null | string>(null)

  const signin = async (apiRoot: string, param: SigninData): Promise<void> => {
    try {
      const response = await fetch(`${apiRoot}/auth/signin`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
        body: JSON.stringify(param),
      })

      if (!response.ok) {
        const errorResponse = await response.json()
        if (!isApiError(errorResponse)) {
          const error = new Error(
            errorResponse.message ??
              `Failed to fetch data from the API. Status: ${response.status}`
          )

          throw error
        }

        setError(errorResponse.userMessage)
      }

      const data = await response.json()
      if (!isUser(data)) {
        const error = new Error(
          `Failed to fetch data from the API. Status: ${response.status}`
        )

        throw error
      }

      setSignedInUser(data)
    } catch (e) {
      console.error(e)

      setError(MESSAGE.UNKNOWN_ERROR)
    }
  }

  return { error, signin }
}
