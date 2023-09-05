import { cookies } from "next/headers"
import { redirect } from "next/navigation"
import { tokenToString } from "typescript"

import { env } from "@/env.mjs"
import { serverFetcher } from "@/lib/fetcher"
import { verifyToken } from "@/lib/verify"
import { SideBar } from "@/components/features/SideBar/SideBar"

interface DefaultLayoutProps {
  children: React.ReactNode
}

export default async function DefaultLayout({ children }: DefaultLayoutProps) {
  // const cookieStore = cookies()
  // const cookie = cookieStore.get("accessToken")

  // if (!cookie) {
  //   console.log("No cookie")
  //   // サインイン画面にリダイレクト
  //   redirect("/signin")
  // }

  // const token = cookie.value

  // try {
  //   const valid = await verifyToken(token)

  //   if (!valid) {
  //     serverFetcher(`${env.API_ROOT}/auth/refresh`, token)
  //   }
  // } catch (e) {
  //   console.log(e)
  //   console.log("Invalid token")

  //   // 不正な場合はサインイン画面にリダイレクト
  //   redirect("/signin")
  // }

  return (
    <div className="fixed top-0 left-0 flex w-full h-screen">
      <main className="flex-1 flex overflow-y-auto relative">
        <div className="ml-64 flex-1 w-128">{children}</div>
        <div className="fixed mid:w-64 midlg:w-96 top-0 right-0 h-screen border-l border-gray-light pl-4">
          <p>開発中(リプライ投稿フォーム)</p>
        </div>
      </main>
      <aside className="w-64 fixed top-0 left-0 h-screen border-r border-gray-light">
        <SideBar className="ml-8 mt-8" />
      </aside>
    </div>
  )
}
