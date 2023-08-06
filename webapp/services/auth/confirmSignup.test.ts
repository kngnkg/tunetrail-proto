import { env } from "@/env.mjs"
import { ApiError } from "@/types/error"

import { ConfirmSignupData, confirmSignup } from "./confirmSignup"

describe("confirmSignup", () => {
  beforeEach(() => {
    jest.resetAllMocks()
  })

  test("ステータスコードが200の場合、nullを返す", async () => {
    const body: ConfirmSignupData = {
      userName: "test",
      code: "test",
    }

    const mock = () =>
      Promise.resolve({
        ok: true,
        status: 200,
        json: () => Promise.resolve(body),
      })
    global.fetch = jest.fn().mockImplementation(mock)

    const data: ConfirmSignupData = {
      userName: "test",
      code: "012345",
    }

    const result = await confirmSignup(env.NEXT_PUBLIC_API_ROOT, data)
    expect(result).toBeNull()
  })

  test("APIからエラーが返ってきた場合、エラーメッセージを返す", async () => {
    const body: ApiError = {
      code: 4101,
      developerMessage: "Invalid confirmation code",
      userMessage: "確認コードが一致しません。",
    }

    const mock = () =>
      Promise.resolve({
        ok: false,
        status: 400,
        json: () => Promise.resolve(body),
      })
    global.fetch = jest.fn().mockImplementation(mock)

    const data: ConfirmSignupData = {
      userName: "test",
      code: "543210",
    }

    const result = await confirmSignup(env.NEXT_PUBLIC_API_ROOT, data)
    expect(result).toBe(body.userMessage)
  })
})
