"use client"

import * as React from "react"

import { env } from "@/env.mjs"
import { User } from "@/types/user"
import { useUserPosts } from "@/hooks/post/use-user-posts"

import { PostList } from "../PostList/PostList"

export interface UserPostListProps
  extends React.HTMLAttributes<HTMLDivElement> {
  user: User
}

export const UserPostList: React.FC<UserPostListProps> = ({
  user,
  className,
  ...props
}) => {
  const { data, error, size, setSize } = useUserPosts(
    env.NEXT_PUBLIC_API_ROOT,
    user.id
  )

  return (
    <>
      <PostList timelines={data} size={size} setSize={setSize} />
    </>
  )
}
