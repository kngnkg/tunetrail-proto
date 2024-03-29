import "@testing-library/jest-dom"
import "@/providers/ToastProvider"
import { render, screen } from "@testing-library/react"

import { Toast, ToastProvider, ToastViewPort } from "./Toast"

describe("ToastViewPort", () => {
  test("positionが正しく適用されること", () => {
    const { rerender } = render(
      <ToastProvider>
        <ToastViewPort />
      </ToastProvider>
    )

    expect(document.querySelector(".top-0")).toBeInTheDocument()

    rerender(
      <ToastProvider>
        <ToastViewPort position="inPlace" />
      </ToastProvider>
    )

    expect(document.querySelector(".top-0")).not.toBeInTheDocument()
  })
})

describe("Toast", () => {
  test("titleとcontentが正しく適用されること", () => {
    const title = "Test Title"
    const content = "Test Content"

    render(
      <ToastProvider>
        <Toast title={title} content={content} />
        <ToastViewPort />
      </ToastProvider>
    )

    expect(screen.getByText(title)).toBeInTheDocument()
    expect(screen.getByText(content)).toBeInTheDocument()
  })

  test("intentが正しく適用されること", () => {
    const { rerender } = render(
      <ToastProvider>
        <Toast intent="success" />
        <ToastViewPort />
      </ToastProvider>
    )

    expect(document.querySelector(".border-green-500")).toBeInTheDocument()

    rerender(
      <ToastProvider>
        <Toast intent="error" />
        <ToastViewPort />
      </ToastProvider>
    )

    expect(document.querySelector(".border-red-500")).toBeInTheDocument()
  })
})
