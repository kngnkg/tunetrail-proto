import { ToastProvider } from "@/providers/ToastProvider"
import { render, screen } from "@testing-library/react"

import { useToast } from "@/hooks/toast/use-toast"

import { Toaster } from "./Toaster"

jest.mock("../../../hooks/toast/use-toast")

const mockUseToast = useToast as jest.MockedFunction<typeof useToast>

describe("Toaster", () => {
  beforeEach(() => {
    mockUseToast.mockReturnValue({
      toasts: [
        { id: "1", intent: "success", description: "test1", open: true },
        { id: "2", intent: "error", description: "test2", open: true },
      ],
      showToast: jest.fn(),
    })
  })

  test("各トーストが適切な内容でレンダリングされること", () => {
    render(<Toaster />, { wrapper: ToastProvider })

    expect(screen.getByText("test1")).toBeInTheDocument()
    expect(screen.getByText("test2")).toBeInTheDocument()
  })

  test("ToastViewPortが正しくレンダリングされること", () => {
    render(<Toaster />, { wrapper: ToastProvider })

    const toastViewPortElement = screen.getByRole("region")
    expect(toastViewPortElement).toBeInTheDocument()
  })
})
