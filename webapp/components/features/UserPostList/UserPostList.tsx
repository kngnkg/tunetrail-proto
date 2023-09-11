"use client"

import * as React from "react"

import { env } from "@/env.mjs"
import { Post } from "@/types/post"
import { Button } from "@/components/ui/Button/Button"

import { PostCard } from "../PostCard/PostCard"

export const UserPostList: React.FC = () => {
  const [posts] = React.useState<Post[]>([])
  const [size, setSize] = React.useState<number>(1)

  return (
    <div className="container mx-auto p-8">
      {posts.map((post: Post, idx: number) => {
        return (
          <div key={idx}>
            <PostCard post={post} className="w-128" />
          </div>
        )
      })}
      <Button onClick={() => setSize(size + 1)}>more</Button>
    </div>
  )
}
