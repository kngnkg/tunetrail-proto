import { render } from "@testing-library/react"

import "@testing-library/jest-dom"
import { Button } from "./Button"

describe("Button", () => {
  test("クラッシュなくレンダリングされること", () => {
    const { getByRole } = render(<Button />)
    expect(getByRole("button")).toBeInTheDocument()
  })

  test("classNameが正しく適用されること", () => {
    const { getByRole } = render(<Button className="test-class" />)
    expect(getByRole("button")).toHaveClass("test-class")
  })

  test("intentが正しく適用されること", () => {
    const { getByRole } = render(<Button intent="primary" />)
    expect(getByRole("button")).toHaveClass("bg-primary-transparent")
  })

  test("sizeが正しく適用されること", () => {
    const { getByRole } = render(<Button size="medium" />)
    expect(getByRole("button")).toHaveClass("h-10")
  })
})
