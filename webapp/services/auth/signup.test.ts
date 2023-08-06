import { env } from "@/env.mjs"
import { ApiError } from "@/types/error"

import { SignupData, signup } from "./signup"

describe("signup", () => {
  beforeEach(() => {
    jest.resetAllMocks()
  })

  test("ステータスコードが200の場合、nullを返す", async () => {
    const body: SignupData = {
      userName: "test",
      name: "test",
      password: "test",
      email: "test",
    }

    const mock = () =>
      Promise.resolve({
        ok: true,
        status: 200,
        json: () => Promise.resolve(body),
      })
    global.fetch = jest.fn().mockImplementation(mock)

    const data: SignupData = {
      userName: "test",
      name: "test",
      password: "test",
      email: "test",
    }

    const result = await signup(env.NEXT_PUBLIC_API_ROOT, data)
    expect(result).toBeNull()
  })

  test("APIからエラーが返ってきた場合、エラーメッセージを返す", async () => {
    const body: ApiError = {
      code: 4202,
      developerMessage: "UserName already entry",
      userMessage: "ユーザー名が既に存在します。",
    }

    const mock = () =>
      Promise.resolve({
        ok: false,
        status: 409,
        json: () => Promise.resolve(body),
      })
    global.fetch = jest.fn().mockImplementation(mock)

    const data: SignupData = {
      userName: "test",
      name: "test",
      password: "test",
      email: "test",
    }

    const result = await signup(env.NEXT_PUBLIC_API_ROOT, data)
    expect(result).toBe(body.userMessage)
  })
})
