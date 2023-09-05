import * as React from "react"
import { HeartFilledIcon, HeartIcon } from "@radix-ui/react-icons"

import { Post } from "@/types/post"
import { mergeClasses } from "@/lib/utils"

export interface LikeButtonProps
  extends React.HTMLAttributes<HTMLButtonElement> {
  post: Post
}

export const LikeButton: React.FC<LikeButtonProps> = ({
  className,
  post,
  ...props
}) => {
  const [isLiked, setIsLiked] = React.useState(false)

  const onClickLike = () => {
    alert("開発中")
    setIsLiked(!isLiked)
  }

  return (
    <>
      {/* {post.isLiked ? ( */}
      {isLiked ? (
        <HeartFilledIcon
          onClick={onClickLike}
          className={mergeClasses(
            "text-primary hover:cursor-pointer",
            className
          )}
        />
      ) : (
        <HeartIcon
          onClick={onClickLike}
          className={mergeClasses(
            "text-gray-lightest hover:text-primary hover:cursor-pointer",
            className
          )}
        />
      )}
    </>
  )
}
