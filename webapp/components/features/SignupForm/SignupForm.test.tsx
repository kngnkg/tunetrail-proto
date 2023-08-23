import { fireEvent, render, screen, waitFor } from "@testing-library/react"
import { useForm } from "react-hook-form"

import { useToast } from "@/hooks/toast/use-toast"

import { SignupForm } from "./SignupForm"

jest.mock("react-hook-form", () => ({
  ...jest.requireActual("react-hook-form"),
  useForm: jest.fn(),
}))

jest.mock("../../../services/auth/signup", () => ({
  isSignupData: jest.fn(),
  signup: jest.fn(),
}))

jest.mock("../../../hooks/toast/use-toast", () => ({
  useToast: jest.fn(),
}))

jest.mock("next/navigation", () => ({
  useRouter: jest.fn(),
}))

describe("SignupForm", () => {
  let mockHandleSubmit: jest.Mock
  let mockRegister: jest.Mock

  beforeEach(() => {
    mockHandleSubmit = jest.fn((cb) => cb)
    mockRegister = jest.fn()
    ;(useForm as jest.Mock).mockReturnValue({
      register: mockRegister,
      handleSubmit: mockHandleSubmit,
      formState: { errors: {} },
    })
    ;(useToast as jest.Mock).mockReturnValue({
      showToast: jest.fn(),
    })
  })

  afterEach(() => {
    jest.clearAllMocks()
  })

  it("クラッシュなくレンダリングされること", () => {
    const { getByText } = render(<SignupForm />)
    expect(getByText("登録")).toBeInTheDocument()
  })

  it("フォームが送信されたらhandleSubmitが呼び出されること", async () => {
    const { getByText } = render(<SignupForm />)

    // フォームに入力
    // Inputコンポーネントのプレースホルダーは内部的にlabelを使用しているため、
    // getByLabelTextで取得する
    fireEvent.change(screen.getByLabelText("アカウント名"), {
      target: { value: "testname" },
    })
    fireEvent.change(screen.getByLabelText("ユーザー名"), {
      target: { value: "testusername" },
    })
    fireEvent.change(screen.getByLabelText("メールアドレス"), {
      target: { value: "test@email.com" },
    })
    fireEvent.change(screen.getByLabelText("パスワード"), {
      target: { value: "Test1234_" },
    })
    fireEvent.change(screen.getByLabelText("パスワード(再度入力)"), {
      target: { value: "Test1234_" },
    })

    const submitButton = getByText("登録")
    fireEvent.submit(submitButton)

    await waitFor(() => {
      expect(mockHandleSubmit).toHaveBeenCalled()
    })
  })
})
