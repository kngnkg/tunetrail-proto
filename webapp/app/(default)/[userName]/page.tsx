import { notFound } from "next/navigation"
import { getUser } from "@/services/user/get-user"

import { env } from "@/env.mjs"
import { FollowButton } from "@/components/features/FollowButton/FollowButton"
import { UserAvatar } from "@/components/features/UserAvatar/UserAvatar"
import { UserPostList } from "@/components/features/UserPostList/UserPostList"
import { getTokensFromCookie } from "@/components/utils"

interface UserPageProps {
  params: { userName: string }
}

// ユーザーページ
export default async function UserPage({ params }: UserPageProps) {
  const tokens = getTokensFromCookie()
  const user = await getUser(env.API_ROOT, tokens, params.userName)

  if (!user) {
    notFound()
  }

  // TODO: ログインユーザーの場合の制御

  return (
    <div className="container mx-auto p-8">
      <div className="flex mb-2">
        <UserAvatar className="h-24 w-24" user={user} />
      </div>
      <div className="flex gap-16 mb-4">
        <div className="flex flex-col gap-0">
          <p className="text-xl">{user.name}</p>
          <p className="text-lg">{user.userName}</p>
        </div>
        <div>
          <FollowButton
            size="medium"
            user={user}
            isFollowing={user.isFollowing}
          />
        </div>
      </div>
      <div>
        <p className="text-xl mb-4">{user.bio}</p>
        <p className="text-xl mb-4">Created: {user.createdAt}</p>
      </div>
      <div>
        <UserPostList user={user} />
      </div>
    </div>
  )
}
