"use client"

import Link from "next/link"

import { env } from "@/env.mjs"
import { Post } from "@/types/post"
import { Timeline, usePosts } from "@/hooks/post/use-posts"

export const PostList: React.FC = () => {
  const { data, error } = usePosts(env.NEXT_PUBLIC_API_ROOT)

  console.log(data)

  return (
    <div className="container mx-auto p-8">
      <Link href="/bob">ボブのページ</Link>
      {data.map((tl: Timeline, tlIdx: number) => {
        return (
          <div key={tlIdx}>
            {tl.posts.map((post: Post, postIdx: number) => {
              return (
                <div key={postIdx}>
                  <p>{post.body}</p>
                </div>
              )
            })}
          </div>
        )
      })}
    </div>
  )
}
