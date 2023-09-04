import Link from "next/link"

import { Post } from "@/types/post"

import { UserAvatar } from "../UserAvatar/UserAvatar"

interface PostCardProps extends React.HTMLAttributes<HTMLDivElement> {
  post: Post
}

export const PostCard: React.FC<PostCardProps> = ({ post, ...props }) => {
  const userPagePath = `/${post.user.userName}`

  return (
    <div
      className="flex items-center justify-between m-2 p-1 border-b border-gray-light"
      {...props}
    >
      <div className="grid gap-1">
        <div className="flex gap-2 items-center">
          <Link href={userPagePath}>
            <UserAvatar user={post.user} />
          </Link>
          <div className="flex flex-col">
            <Link href={userPagePath}>{post.user.name}</Link>
            <Link
              href={userPagePath}
              className="text-sm text-gray-lightest"
            >{`@${post.user.userName}`}</Link>
          </div>
        </div>
        <div className="mt-3 mb-2 ml-12">
          <p>{post.body}</p>
        </div>
      </div>
    </div>
  )
}
