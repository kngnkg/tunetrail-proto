import * as React from "react"
import { ChatBubbleIcon } from "@radix-ui/react-icons"

import { Post } from "@/types/post"
import { mergeClasses } from "@/lib/utils"

import { NewPostDialog } from "../NewPostDialog/NewPostDialog"

export interface ReplyButtonProps
  extends React.HTMLAttributes<HTMLButtonElement> {
  post: Post
}

export const ReplyButton: React.FC<ReplyButtonProps> = ({
  className,
  post,
  ...props
}) => {
  return (
    <div>
      <NewPostDialog parentPost={post}>
        <ChatBubbleIcon
          className={mergeClasses(
            "text-gray-lightest hover:cursor-pointer hover:text-primary",
            className
          )}
        />
      </NewPostDialog>
    </div>
  )
}
