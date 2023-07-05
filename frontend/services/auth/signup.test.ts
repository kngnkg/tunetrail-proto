import { MESSAGE } from "@/config/messages"

import { signup } from "./signup"

describe("signup", () => {
  const context = {
    apiRoot: "ttp://localhost:3000",
  }

  beforeEach(() => {
    jest.resetAllMocks()
  })

  //   test("リクエストのフォーマットが正しいこと", () => {})

  test("ステータスコードが200の場合、nullを返す", async () => {
    const body = {
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

    const data = {
      userName: "test",
      name: "test",
      password: "test",
      email: "test",
    }

    const result = await signup(context, data)
    expect(result).toBeNull()
  })

  describe("ステータスコードが400の場合", () => {
    test("ユーザーネームが重複していた場合", async () => {
      const mock = () =>
        Promise.resolve({
          ok: false,
          status: 400,
          json: () =>
            Promise.resolve({ msg: "ユーザー名が既に登録されています。" }),
        })
      global.fetch = jest.fn().mockImplementation(mock)

      const data = {
        userName: "test",
        name: "test",
        password: "test",
        email: "test",
      }

      const result = await signup(context, data)
      expect(result).toBe(MESSAGE.DUP_USERNAME)
    })

    test("メールアドレスが重複していた場合", async () => {
      const mock = () =>
        Promise.resolve({
          ok: false,
          status: 400,
          json: () =>
            Promise.resolve({ msg: "メールアドレスが既に登録されています。" }),
        })
      global.fetch = jest.fn().mockImplementation(mock)

      const data = {
        userName: "test",
        name: "test",
        password: "test",
        email: "test",
      }

      const result = await signup(context, data)
      expect(result).toBe(MESSAGE.DUP_EMAIL)
    })
  })

  test("ステータスコードが500の場合", async () => {
    const mock = () =>
      Promise.resolve({
        ok: false,
        status: 500,
        json: () =>
          Promise.resolve({ msg: "サーバー内部でエラーが発生しました。" }),
      })
    global.fetch = jest.fn().mockImplementation(mock)

    const data = {
      userName: "test",
      name: "test",
      password: "test",
      email: "test",
    }

    const result = await signup(context, data)
    expect(result).toBe(MESSAGE.UNKNOWN_ERROR)
  })
})
