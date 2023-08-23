import { notFound } from "next/navigation"
import { getUser } from "@/services/user/get-user"

import { env } from "@/env.mjs"

interface UserPageProps {
  params: { userName: string }
}

// ユーザーページ
export default async function UserPage({ params }: UserPageProps) {
  const user = await getUser(env.API_ROOT, params.userName)

  if (!user) {
    notFound()
  }

  return (
    <div className="container mx-auto p-8">
      <h1 className="text-3xl mb-8">User Page</h1>
      <div>
        <p className="text-xl mb-4">User Name: {user.userName}</p>
        <p className="text-xl mb-4">Name: {user.name}</p>
        <p className="text-xl mb-4">IconUrl: {user.iconUrl}</p>
        <p className="text-xl mb-4">Bio: {user.bio}</p>
        <p className="text-xl mb-4">Created: {user.createdAt}</p>
      </div>
    </div>
  )
}
