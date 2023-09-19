"use client"

import { env } from "@/env.mjs"
import { useTimeline } from "@/hooks/post/use-timeline"

import { PostList } from "../PostList/PostList"

export const TimeLine: React.FC = () => {
  const { data, error, size, setSize, mutatePost } = useTimeline(
    env.NEXT_PUBLIC_API_ROOT
  )

  return (
    <>
      <PostList
        timelines={data}
        mutatePost={mutatePost}
        size={size}
        setSize={setSize}
      />
    </>
  )
}
