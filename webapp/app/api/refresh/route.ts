import { NextRequest, NextResponse } from "next/server"

import { env } from "@/env.mjs"
import fetcher from "@/lib/fetcher"

type RefreshToken = {
  IdToken: string
  RefreshToken: string
}

export async function POST(req: NextRequest) {
  const idCookie = req.cookies.get("idToken")

  if (!idCookie) {
    console.log("idToken is not found")
    return new Response("Unauthorized", { status: 403 })
  }

  const refreshCookie = req.cookies.get("refreshToken")

  if (!refreshCookie) {
    console.log("refreshToken is not found")
    return new Response("Unauthorized", { status: 403 })
  }

  const body: RefreshToken = {
    IdToken: idCookie.value,
    RefreshToken: refreshCookie.value,
  }

  const apiResp = await fetcher(`${env.API_ROOT}/auth/refresh`, {
    method: "POST",
    body: JSON.stringify(body),
  })

  let resp = new NextResponse()

  resp.cookies.set("idToken", apiResp.idToken, {
    httpOnly: true,
    maxAge: 60 * 60 * 24 * 7, // 考える
    path: "/",
    domain: `.${env.ALLOWED_DOMAIN}`,
    sameSite: "none",
    secure: true,
  })

  resp.cookies.set("accessToken", apiResp.accessToken, {
    httpOnly: true,
    maxAge: 60 * 60 * 24 * 7, // 考える
    path: "/",
    domain: `.${env.ALLOWED_DOMAIN}`,
    sameSite: "none",
    secure: true,
  })

  return resp
}
