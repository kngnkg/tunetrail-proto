import React from "react"
import { Health } from "@/types"
import { render, screen } from "@testing-library/react"

import "@testing-library/jest-dom"
import HealthStatus from "./HealthStatus"

describe("HealthStatus コンポーネント", () => {
  const health: Health = {
    health: "green",
    database: "red",
  }

  test("ヘルスチェックのステータスに基づいてHealth Checkのインジケータを表示する", () => {
    render(<HealthStatus health={health} />)

    const healthIndicator = screen.getByText("GREEN")
    expect(healthIndicator).toBeInTheDocument()
    expect(healthIndicator).toHaveClass("text-green-500")
  })

  test("データベースのステータスに基づいてDatabase Checkのインジケータを表示する", () => {
    render(<HealthStatus health={health} />)

    const databaseIndicator = screen.getByText("RED")
    expect(databaseIndicator).toBeInTheDocument()
    expect(databaseIndicator).toHaveClass("text-red-500")
  })
})
