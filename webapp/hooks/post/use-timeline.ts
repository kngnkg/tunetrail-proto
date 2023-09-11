import useSWRInfinite from "swr/infinite"

import { Timeline } from "@/types/post"
import { clientFetcher } from "@/lib/fetcher"

export interface PostParams {
  body: string
}

export const useTimeline = (
  apiRoot: string
): {
  data: Timeline[]
  size: number
  setSize: (
    size: number | ((size: number) => number)
  ) => Promise<any[] | undefined>
  error: null | string
  addPost: (param: PostParams) => Promise<void>
} => {
  const getKey = (pageIndex: number, previousPageData: Timeline) => {
    // 最後に到達した場合
    if (previousPageData && previousPageData.pagination.nextCursor === "") {
      return null
    }

    // 最初のページでは、`previousPageData` がない
    if (pageIndex === 0) {
      return `${apiRoot}/users/timelines`
    }

    // API のエンドポイントにカーソルを追加する
    return `${apiRoot}/users/timelines?next_cursor=${previousPageData.pagination.nextCursor}`
  }

  const { data, error, isLoading, isValidating, mutate, size, setSize } =
    useSWRInfinite(getKey, clientFetcher)

  const addPost = async (param: PostParams): Promise<void> => {
    try {
      const res = await clientFetcher(`${apiRoot}/posts`, {
        method: "POST",
        body: JSON.stringify({
          body: param.body,
        }),
      })

      if (!data) {
        mutate(
          [
            {
              posts: [res],
              pagination: {
                nextCursor: "",
                previousCursor: "",
                limit: 0,
              },
            },
          ],
          false
        )
        return
      }

      const pagination = data[0].pagination
      const tl: Timeline = {
        posts: [res],
        pagination,
      }

      mutate([tl, ...data], false)
    } catch (e) {
      throw e
    }
  }

  if (isLoading) return { data: [], error: null, size, setSize, addPost }

  if (error) return { data: [], error: error.message, size, setSize, addPost }

  if (!data) return { data: [], error: null, size, setSize, addPost }

  return { data, error, size, setSize, addPost }
}
