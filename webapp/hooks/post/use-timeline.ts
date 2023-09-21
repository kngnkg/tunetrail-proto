import useSWRInfinite from "swr/infinite"

import { Post, Timeline } from "@/types/post"
import { clientFetcher } from "@/lib/fetcher"

const fetcher = async (
  resource: RequestInfo,
  init?: RequestInit
): Promise<Timeline> => {
  try {
    const res = await clientFetcher(resource, init)

    const posts: Post[] = res.posts.map((post: Post) => {
      return {
        ...post,
        createdAt: new Date(post.createdAt),
      }
    })

    const timeline: Timeline = {
      posts,
      pagination: res.pagination,
    }

    return timeline
  } catch (e) {
    throw e
  }
}

export interface PostParams {
  body: string
}

// TODO: エンドポイントを引数に取って再利用できるようにする
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
  mutatePost: (post: Post) => void
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
    useSWRInfinite(getKey, fetcher)

  const mutatePost = (post: Post): void => {
    if (!data) return

    const updated: Timeline[] = data.map((tl) => {
      const posts = tl.posts.map((p) => {
        if (p.id === post.id) {
          return post
        }

        return p
      })

      return {
        ...tl,
        posts,
      }
    })

    mutate(updated, false)
  }

  const addPost = async (param: PostParams): Promise<void> => {
    try {
      const res = await clientFetcher(`${apiRoot}/posts`, {
        method: "POST",
        body: JSON.stringify({
          body: param.body,
        }),
      })

      const post: Post = {
        ...res,
        createdAt: new Date(res.createdAt),
      }

      if (!data) {
        const tl: Timeline = {
          posts: [post],
          pagination: {
            nextCursor: "",
            previousCursor: "",
            limit: 0,
          },
        }

        mutate([tl], false)
        return
      }

      const pagination = data[0].pagination
      const tl: Timeline = {
        posts: [post],
        pagination,
      }

      mutate([tl, ...data], false)
    } catch (e) {
      throw e
    }
  }

  if (isLoading)
    return { data: [], error: null, size, setSize, addPost, mutatePost }

  if (error)
    return {
      data: [],
      error: error.message,
      size,
      setSize,
      addPost,
      mutatePost,
    }

  if (!data)
    return { data: [], error: null, size, setSize, addPost, mutatePost }

  return { data, error, size, setSize, addPost, mutatePost }
}
