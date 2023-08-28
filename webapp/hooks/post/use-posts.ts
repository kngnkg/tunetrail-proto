import * as React from "react"
import useSWRInfinite from "swr/infinite"

import { Post } from "@/types/post"
import { clientFetcher } from "@/lib/fetcher"

export type Timeline = {
  posts: Post[]
  pagenation: {
    nextCursor: string
    previousCursor: string
    limit: number
  }
}

export interface PostParams {
  body: string
}

export const usePosts = (
  apiRoot: string
): {
  // data: Post[]
  data: Timeline[]
  error: null | string
  addPost: (param: PostParams) => Promise<void>
} => {
  const getKey = (pageIndex: number, previousPageData: Timeline) => {
    // 最後に到達した場合
    if (previousPageData && !previousPageData.posts) return null

    // 最初のページでは、`previousPageData` がない
    if (pageIndex === 0) return `${apiRoot}/users/timelines`

    // API のエンドポイントにカーソルを追加する
    // return `${apiRoot}/posts?cursor=${previousPageData.pagenation.nextCursor}`
    return `${apiRoot}/users/timelines`
  }

  const { data, error, isLoading, isValidating, mutate, size, setSize } =
    useSWRInfinite(getKey, clientFetcher)
  // const [nextCursor, setNextCursor] = React.useState<string | null>(null)

  const addPost = async (param: PostParams): Promise<void> => {
    try {
      const res = await clientFetcher(`${apiRoot}/posts`, {
        method: "POST",
        body: JSON.stringify({
          body: param.body,
        }),
      })

      // setPosts(res)
    } catch (e) {
      if (e instanceof Error) {
        // setError(e.message)
        return
      }

      // setError(MESSAGE.UNKNOWN_ERROR)
    }
  }

  if (isLoading) return { data: [], error: null, addPost }

  if (error) return { data: [], error: error.message, addPost }

  if (!data) return { data: [], error: null, addPost }

  return { data, error, addPost }
}
