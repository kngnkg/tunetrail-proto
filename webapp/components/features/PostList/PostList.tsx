"use client"

import { Post, Timeline } from "@/types/post"
import { Button } from "@/components/ui/Button/Button"

import { PostCard } from "../PostCard/PostCard"

export interface PostListProps extends React.HTMLAttributes<HTMLDivElement> {
  timelines: Timeline[]
  size: number
  setSize: (
    size: number | ((size: number) => number)
  ) => Promise<any[] | undefined>
}

export const PostList: React.FC<PostListProps> = ({
  timelines,
  size,
  setSize,
  ...props
}) => {
  return (
    <div className="container mx-auto p-8">
      {timelines.map((tl: Timeline, tlIdx: number) => {
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
