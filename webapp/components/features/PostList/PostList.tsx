"use client"

import { env } from "@/env.mjs"
import { Post, Timeline } from "@/types/post"
import { usePosts } from "@/hooks/post/use-posts"
import { Button } from "@/components/ui/Button/Button"

import { PostCard } from "../PostCard/PostCard"

export const PostList: React.FC = () => {
  const { data, error, size, setSize } = usePosts(env.NEXT_PUBLIC_API_ROOT)

  return (
    <div className="container mx-auto p-8">
      {data.map((tl: Timeline, tlIdx: number) => {
        return (
          <div key={tlIdx}>
            {tl.posts.map((post: Post, postIdx: number) => {
              return (
                <div key={postIdx}>
                  <PostCard post={post} className="w-128" />
                </div>
              )
            })}
          </div>
        )
      })}
      <Button onClick={() => setSize(size + 1)}>more</Button>
    </div>
  )
}
