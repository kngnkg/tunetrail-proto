import useSWRInfinite from "swr/infinite"

import { Timeline } from "@/types/post"
import { clientFetcher } from "@/lib/fetcher"

export const useUserPosts = (apiRoot: string, userId: string) => {
  const getKey = (pageIndex: number, previousPageData: Timeline) => {
    // 最後に到達した場合
    if (previousPageData && previousPageData.pagination.nextCursor === "") {
      return null
    }

    // 最初のページでは、`previousPageData` がない
    if (pageIndex === 0) {
      return `${apiRoot}/users/${userId}/posts`
    }

    // API のエンドポイントにカーソルを追加する
    return `${apiRoot}/users/${userId}/posts?next_cursor=${previousPageData.pagination.nextCursor}`
  }

  const { data, error, isLoading, isValidating, mutate, size, setSize } =
    useSWRInfinite(getKey, clientFetcher)

  if (isLoading) return { data: [], error: null, size, setSize }

  if (error) return { data: [], error: error.message, size, setSize }

  if (!data) return { data: [], error: null, size, setSize }

  return { data, error, size, setSize }
}
