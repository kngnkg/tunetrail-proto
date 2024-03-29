import * as React from "react"

import { MESSAGE } from "@/config/messages"
import { clientFetcher } from "@/lib/fetcher"

export interface UseFollowProps {
  apiRoot: string
  isFollowing: boolean
}

export const useFollow = ({
  ...props
}: UseFollowProps): {
  isFollowing: boolean
  error: null | string
  follow: (signinUserName: string, followUserName: string) => Promise<void>
  unfollow: (signinUserName: string, followUserName: string) => Promise<void>
} => {
  const [isFollowing, setIsFollowing] = React.useState(props.isFollowing)
  const [error, setError] = React.useState<null | string>(null)

  const follow = async (
    signinUserId: string,
    followUserId: string
  ): Promise<void> => {
    try {
      const resp = await clientFetcher(
        `${props.apiRoot}/users/${signinUserId}/follow`,
        {
          method: "POST",
          body: JSON.stringify({
            followee_user_id: followUserId,
          }),
        }
      )

      setIsFollowing(true)
    } catch (e) {
      if (e instanceof Error) {
        setError(e.message)
        return
      }

      setError(MESSAGE.UNKNOWN_ERROR)
    }
  }

  const unfollow = async (
    signinUserId: string,
    followUserId: string
  ): Promise<void> => {
    try {
      const resp = await clientFetcher(
        `${props.apiRoot}/users/${signinUserId}/follow`,
        {
          method: "DELETE",
          body: JSON.stringify({
            followee_user_id: followUserId,
          }),
        }
      )

      setIsFollowing(false)
    } catch (e) {
      if (e instanceof Error) {
        setError(e.message)
        return
      }

      setError(MESSAGE.UNKNOWN_ERROR)
    }
  }

  return {
    isFollowing,
    error,
    follow,
    unfollow,
  }
}
