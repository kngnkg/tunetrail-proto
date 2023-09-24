import { Post } from "@/types/post"
import { Tokens } from "@/types/tokens"
import { serverFetcher } from "@/lib/fetcher"

export const getPost = async (
  apiRoot: string,
  tokens: Tokens,
  postId: string
): Promise<Post | null> => {
  try {
    const res = await serverFetcher(`${apiRoot}/posts/${postId}`, {
      cache: "no-store",
      headers: {
        Cookie: `idToken=${tokens.idToken}; accessToken=${tokens.accessToken};`,
      },
    })

    const post: Post = {
      ...res,
      createdAt: new Date(res.createdAt),
    }

    return post
  } catch (e) {
    throw e
  }
}
