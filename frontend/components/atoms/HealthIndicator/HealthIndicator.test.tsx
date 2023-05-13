import React from "react"
import { render, screen } from "@testing-library/react"

import "@testing-library/jest-dom"
import HealthIndicator from "./HealthIndicator"

describe("HealthIndicator コンポーネント", () => {
  test("statusがgreenの場合、GREENを緑色で表示する", () => {
    render(<HealthIndicator status="green" />)

    // 指定したテキストが表示されていることを確認する
    const statusElement = screen.getByText("GREEN")
    expect(statusElement).toBeInTheDocument()
    // 指定したクラスが付与されていることを確認する
    expect(statusElement).toHaveClass("text-green-500")
  })

  test("statusがorangeの場合、ORANGEを赤色で表示する", () => {
    render(<HealthIndicator status="orange" />)

    const statusElement = screen.getByText("ORANGE")
    expect(statusElement).toBeInTheDocument()
    expect(statusElement).toHaveClass("text-red-500")
  })

  test("statusがredの場合、REDを赤色で表示する", () => {
    render(<HealthIndicator status="red" />)

    const statusElement = screen.getByText("RED")
    expect(statusElement).toBeInTheDocument()
    expect(statusElement).toHaveClass("text-red-500")
  })
})
