import { cookies } from "next/headers"
import { redirect } from "next/navigation"

import { Tokens } from "@/types/tokens"

export function getTokensFromCookie(): Tokens {
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

  return tokens
}
