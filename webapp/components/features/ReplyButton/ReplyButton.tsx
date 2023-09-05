import * as React from "react"
import { ChatBubbleIcon } from "@radix-ui/react-icons"

import { Post } from "@/types/post"
import { mergeClasses } from "@/lib/utils"

export interface ReplyButtonProps
  extends React.HTMLAttributes<HTMLButtonElement> {
  post: Post
}

export const ReplyButton: React.FC<ReplyButtonProps> = ({
  className,
  post,
  ...props
}) => {
  const [clicked, setClicked] = React.useState(false)

  const onClick = () => {
    alert("開発中")
    setClicked(!clicked)
  }

  return (
    <ChatBubbleIcon
      onClick={onClick}
      className={mergeClasses(
        "text-gray-lightest hover:cursor-pointer hover:text-primary",
        className
      )}
    />
  )
}
