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
  //   liked: boolean
  error: null | string
  addLike: () => Promise<void>
  deleteLike: () => Promise<void>
} => {
  //   const [liked, setLiked] = React.useState(props.post.liked)
  const [error, setError] = React.useState<null | string>(null)

  const addLike = async (): Promise<void> => {
    try {
      const resp = await clientFetcher(
        `${props.apiRoot}/posts/${props.post.id}/likes`,
        {
          method: "POST",
        }
      )

      //   setLiked(true)
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

      //   setLiked(false)
    } catch (e) {
      if (e instanceof Error) {
        setError(e.message)
        return
      }

      setError(MESSAGE.UNKNOWN_ERROR)
    }
  }

  return {
    // liked,
    error,
    addLike,
    deleteLike,
  }
}
