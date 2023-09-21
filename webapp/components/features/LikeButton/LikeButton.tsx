import * as React from "react"
import { HeartFilledIcon, HeartIcon } from "@radix-ui/react-icons"

import { env } from "@/env.mjs"
import { Post } from "@/types/post"
import { mergeClasses } from "@/lib/utils"
import { useSignedInUser } from "@/hooks/auth/use-signedin-user"
import { useLike } from "@/hooks/like/use-like"
import { useTimeline } from "@/hooks/post/use-timeline"
import { useToast } from "@/hooks/toast/use-toast"

export interface LikeButtonProps
  extends React.HTMLAttributes<HTMLButtonElement> {
  post: Post
}

export const LikeButton: React.FC<LikeButtonProps> = ({
  className,
  post,
  ...props
}) => {
  const signedInUser = useSignedInUser()
  const { showToast } = useToast()
  const { mutatePost } = useTimeline(env.NEXT_PUBLIC_API_ROOT)
  const { error, addLike, deleteLike } = useLike({
    apiRoot: env.NEXT_PUBLIC_API_ROOT,
    post,
  })

  const addLikeInternal = async () => {
    await addLike()

    if (error) {
      showToast({
        intent: "error",
        description: error,
      })
    }

    const updated: Post = {
      ...post,
      liked: true,
      likesCount: post.likesCount + 1,
    }

    mutatePost(updated)
  }

  const deleteLikeInternal = async () => {
    await deleteLike()

    if (error) {
      showToast({
        intent: "error",
        description: error,
      })
    }

    const updated: Post = {
      ...post,
      liked: false,
      likesCount: post.likesCount - 1,
    }

    mutatePost(updated)
  }

  const onClickToggleLike = async () => {
    if (!signedInUser) return

    if (post.liked) {
      await deleteLikeInternal()
      return
    }

    await addLikeInternal()
  }

  return (
    <div className="flex gap-1 items-center">
      {post.liked ? (
        <HeartFilledIcon
          onClick={onClickToggleLike}
          className={mergeClasses(
            "text-primary hover:cursor-pointer",
            className
          )}
        />
      ) : (
        <HeartIcon
          onClick={onClickToggleLike}
          className={mergeClasses(
            "text-gray-lightest hover:text-primary hover:cursor-pointer",
            className
          )}
        />
      )}
      {post.likesCount > 0 ? (
        <span className="text-gray-lightest">{post.likesCount}</span>
      ) : null}
    </div>
  )
}
