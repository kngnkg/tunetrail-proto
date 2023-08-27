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
    <div className="flex min-h-screen">
      <main className="flex-1">{children}</main>
      <aside className="w-1/3">
        <SideBar />
      </aside>
    </div>
  )
}
