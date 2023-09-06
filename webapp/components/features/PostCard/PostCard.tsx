import Link from "next/link"
import { ChatBubbleIcon } from "@radix-ui/react-icons"

import { Post } from "@/types/post"
import { mergeClasses } from "@/lib/utils"
import {
  Card,
  CardContent,
  CardHeader,
  CardHooter,
} from "@/components/ui/Card/Card"

import { LikeButton } from "../LikeButton/LikeButton"
import { ReplyButton } from "../ReplyButton/ReplyButton"
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
  const postPagePath = `/${post.user.userName}/${post.id}`

  const hooterIconClasses = "w-5 h-5"

  return (
    <Card className={mergeClasses("flex gap-2 pt-2", className)} {...props}>
      <CardHeader className="w-12 h-full pt-2 pl-2 flex-shrink-0">
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
        <div className="pb-2 pl-2 pr-3">
          <Link href={postPagePath}>
            <p>{post.body}</p>
          </Link>
        </div>
        <CardHooter className="flex gap-6 items-center pb-1">
          <ReplyButton post={post} className={hooterIconClasses} />
          <LikeButton post={post} className={hooterIconClasses} />
        </CardHooter>
      </CardContent>
    </Card>
  )
}
