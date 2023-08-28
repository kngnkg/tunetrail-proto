import * as React from "react"

import { Post } from "@/types/post"
import { MESSAGE } from "@/config/messages"
import { clientFetcher } from "@/lib/fetcher"

export interface PostParams {
  body: string
}

export const usePosts = (
  apiRoot: string
): {
  posts: Post[]
  error: null | string
  addPost: (param: PostParams) => Promise<void>
} => {
  const [posts, setPosts] = React.useState<Post[]>([])
  const [error, setError] = React.useState<null | string>(null)

  // const fetchPosts = async (): Promise<void> => {

  const addPost = async (param: PostParams): Promise<void> => {
    try {
      const res = await clientFetcher(`${apiRoot}/posts`, {
        method: "POST",
        body: JSON.stringify({
          body: param.body,
        }),
      })

      setPosts(res)
    } catch (e) {
      if (e instanceof Error) {
        setError(e.message)
        return
      }

      setError(MESSAGE.UNKNOWN_ERROR)
    }
  }

  return { posts, error, addPost }
}
