import { cookies } from "next/headers"
import Link from "next/link"
import { notFound, redirect } from "next/navigation"
import { getUser } from "@/services/user/get-user"

import { env } from "@/env.mjs"
import { Tokens } from "@/types/tokens"
import { FollowButton } from "@/components/features/FollowButton/FollowButton"

interface UserPageProps {
  params: { userName: string }
}

// ユーザーページ
export default async function UserPage({ params }: UserPageProps) {
  const cookieStore = cookies()
  const idCookie = cookieStore.get("idToken")
  const accessCookie = cookieStore.get("accessToken")

  if (idCookie === undefined || accessCookie === undefined) {
    // 暫定としてサインインページにリダイレクト
    redirect("/signin")
  }

  const tokens: Tokens = {
    idToken: idCookie?.value,
    accessToken: accessCookie?.value,
    refreshToken: "",
  }

  const user = await getUser(env.API_ROOT, tokens, params.userName)

  if (!user) {
    notFound()
  }

  // TODO: ログインユーザーの場合の制御

  return (
    <div className="container mx-auto p-8">
      <h1 className="text-3xl mb-8">User Page</h1>
      <div>
        <FollowButton userName={user.userName} isFollowing={user.isFollowing} />
      </div>
      <div>
        <p className="text-xl mb-4">User Name: {user.userName}</p>
        <p className="text-xl mb-4">Name: {user.name}</p>
        <p className="text-xl mb-4">IconUrl: {user.iconUrl}</p>
        <p className="text-xl mb-4">Bio: {user.bio}</p>
        <p className="text-xl mb-4">IsFollowing: {user.isFollowing}</p>
        <p className="text-xl mb-4">IsFollowed: {user.isFollowed}</p>
        <p className="text-xl mb-4">Created: {user.createdAt}</p>
      </div>
    </div>
  )
}
