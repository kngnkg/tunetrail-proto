import { render } from "@testing-library/react"

import { Label } from "./Label"

describe("Label", () => {
  test("クラッシュなくレンダリングされること", () => {
    const { getByText } = render(<Label>Test Text</Label>)
    expect(getByText("Test Text")).toBeInTheDocument()
  })

  test("classNameが正しく適用されること", () => {
    const { getByText } = render(
      <Label className="test-class">Test Text</Label>
    )
    expect(getByText("Test Text")).toHaveClass("test-class")
  })
})
