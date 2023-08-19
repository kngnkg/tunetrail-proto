import { NextRequest, NextResponse } from "next/server"

import { env } from "@/env.mjs"
import fetcher from "@/lib/fetcher"

export async function POST(req: NextRequest) {
  const refreshToken = req.cookies.get("refreshToken")

  const apiResp = await fetcher(`${env.NEXT_PUBLIC_API_ROOT}/auth/refresh`, {
    method: "POST",
    // credentials: "include",
    headers: {
      Cookie: `refreshToken=${refreshToken}`,
    },
  })

  let resp = NextResponse.next()

  resp.cookies.set("accessToken", apiResp.accessToken, {
    httpOnly: true,
    maxAge: 60 * 60 * 24 * 7,
    path: "/",
    sameSite: "none",
    secure: true,
  })

  return resp
}
