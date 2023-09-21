import * as React from "react"

import { Post } from "@/types/post"
import { MESSAGE } from "@/config/messages"
import { clientFetcher } from "@/lib/fetcher"

export interface UseLikeProps {
  apiRoot: string
  post: Post
}

export const useLike = ({
  ...props
}: UseLikeProps): {
  error: null | string
  addLike: () => Promise<void>
  deleteLike: () => Promise<void>
} => {
  const [error, setError] = React.useState<null | string>(null)

  const addLike = async (): Promise<void> => {
    try {
      const resp = await clientFetcher(
        `${props.apiRoot}/posts/${props.post.id}/likes`,
        {
          method: "POST",
        }
      )
    } catch (e) {
      if (e instanceof Error) {
        setError(e.message)
        return
      }

      setError(MESSAGE.UNKNOWN_ERROR)
    }
  }

  const deleteLike = async (): Promise<void> => {
    try {
      const resp = await clientFetcher(
        `${props.apiRoot}/posts/${props.post.id}/likes`,
        {
          method: "DELETE",
        }
      )
    } catch (e) {
      if (e instanceof Error) {
        setError(e.message)
        return
      }

      setError(MESSAGE.UNKNOWN_ERROR)
    }
  }

  return {
    error,
    addLike,
    deleteLike,
  }
}
