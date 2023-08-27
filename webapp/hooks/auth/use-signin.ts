import * as React from "react"

import { isUser } from "@/types/user"
import { MESSAGE } from "@/config/messages"
import { clientFetcher } from "@/lib/fetcher"

import { useSetSignedInUser } from "./use-signedin-user"

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
  const setSignedInUser = useSetSignedInUser()
  const [error, setError] = React.useState<null | string>(null)

  const signin = async (apiRoot: string, param: SigninData): Promise<void> => {
    try {
      const data = await clientFetcher(`${apiRoot}/auth/signin`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(param),
      })

      if (!isUser(data)) {
        throw new Error("Failed to fetch data from the API.")
      }

      setSignedInUser(data)
    } catch (e) {
      console.error(e)
      if (e instanceof Error) {
        setError(e.message)
        return
      }

      setError(MESSAGE.UNKNOWN_ERROR)
    }
  }

  return { error, signin }
}
