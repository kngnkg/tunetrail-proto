import { notFound } from "next/navigation"
import { getPost } from "@/services/post/get-post"

import { env } from "@/env.mjs"
import { PostCard } from "@/components/features/PostCard/PostCard"
import { getTokensFromCookie } from "@/components/utils"

interface PostPageProps {
  params: { postId: string }
}

export default async function PostPage({ params }: PostPageProps) {
  const tokens = getTokensFromCookie()

  const post = await getPost(env.API_ROOT, tokens, params.postId)

  if (!post) {
    notFound()
  }

  return (
    <div className="container mx-auto p-8">
      <PostCard post={post} />
    </div>
  )
}
