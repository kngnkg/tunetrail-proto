import { fireEvent, render } from "@testing-library/react"

import "@testing-library/jest-dom"
import { Input } from "./Input"

describe("Input", () => {
  test("クラッシュなくレンダリングされること", () => {
    const { getByRole } = render(<Input />)
    expect(getByRole("textbox")).toBeInTheDocument()
  })

  test("classNameが正しく適用されること", () => {
    const { getByRole } = render(<Input className="test-class" />)
    expect(getByRole("textbox")).toHaveClass("test-class")
  })

  test("プレースホルダーが正しく適用されること", () => {
    const { queryByText } = render(<Input placeholderText="Test Placeholder" />)
    expect(queryByText("Test Placeholder")).toBeInTheDocument()
  })

  test("テキストが入力された場合にonChangeイベントが発火すること", () => {
    const handleChange = jest.fn()
    const { getByRole } = render(<Input onChange={handleChange} />)
    const input = getByRole("textbox")

    fireEvent.change(input, { target: { value: "Test input" } })

    expect(handleChange).toHaveBeenCalledTimes(1)
  })

  test("フォーカスとブラー時の状態変化をテストする", () => {
    const { getByRole, getByText } = render(
      <Input placeholderText="Test Placeholder" />
    )
    const input = getByRole("textbox")
    const placeholder = getByText("Test Placeholder")

    fireEvent.focus(input)
    expect(placeholder).toHaveClass("text-tiny")

    fireEvent.blur(input)
    expect(placeholder).toHaveClass("text-sm")
  })
})
