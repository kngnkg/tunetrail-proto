import Link from "next/link"
import { ChatBubbleIcon, HeartIcon } from "@radix-ui/react-icons"

import { Post } from "@/types/post"
import { mergeClasses } from "@/lib/utils"
import {
  Card,
  CardContent,
  CardHeader,
  CardHooter,
} from "@/components/ui/Card/Card"

import { UserAvatar } from "../UserAvatar/UserAvatar"

interface PostCardProps extends React.HTMLAttributes<HTMLDivElement> {
  post: Post
  className?: string
}

export const PostCard: React.FC<PostCardProps> = ({
  post,
  className,
  ...props
}) => {
  const userPagePath = `/${post.user.userName}`

  const hooterIconClasses = "w-4 h-4 text-gray-lightest"

  return (
    <Card className={mergeClasses("flex gap-2 pt-2", className)} {...props}>
      <CardHeader className="w-12 h-full pt-1 pl-2 flex-shrink-0">
        <Link href={userPagePath}>
          <UserAvatar user={post.user} />
        </Link>
      </CardHeader>
      <CardContent className="flex flex-col gap-2">
        <div className="flex flex-col gap-0">
          <Link href={userPagePath}>{post.user.name}</Link>
          <Link
            href={userPagePath}
            className="text-sm text-gray-lightest"
          >{`@${post.user.userName}`}</Link>
        </div>
        <div className="pb-1 pl-2 pr-3">
          <p>{post.body}</p>
        </div>
        <CardHooter className="flex gap-4 items-center pb-1">
          <ChatBubbleIcon className={hooterIconClasses} />
          <HeartIcon className={hooterIconClasses} />
        </CardHooter>
      </CardContent>
    </Card>
  )
}
