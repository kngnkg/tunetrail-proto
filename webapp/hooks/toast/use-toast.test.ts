import "@testing-library/jest-dom"
import { ToastProvider } from "@/providers/ToastProvider"
import { act, renderHook } from "@testing-library/react"

import { useToast } from "./use-toast"

jest.useFakeTimers()

describe("useToast", () => {
  test("showToastが呼ばれたときにトーストが表示されること", () => {
    const { result } = renderHook(() => useToast(), { wrapper: ToastProvider })

    act(() => {
      result.current.showToast({ intent: "default", description: "test" })
    })

    expect(result.current.toasts).toHaveLength(1)
    const toast = result.current.toasts[0]
    expect(toast).toMatchObject({
      intent: "default",
      description: "test",
      open: true,
      id: expect.any(String),
    })
  })

  test("トーストに一意のIDが付与されること", () => {
    const { result } = renderHook(() => useToast(), { wrapper: ToastProvider })

    act(() => {
      result.current.showToast({ intent: "default", description: "test 1" })
      result.current.showToast({ intent: "default", description: "test 2" })
    })

    expect(result.current.toasts).toHaveLength(2)
    expect(result.current.toasts[0].id).not.toEqual(result.current.toasts[1].id)
  })

  test("TIMEOUTの時間が経過したときにトーストが非表示になること", () => {
    const { result } = renderHook(() => useToast(), { wrapper: ToastProvider })

    act(() => {
      result.current.showToast({ intent: "default", description: "test" })
    })

    expect(result.current.toasts).toHaveLength(1)

    act(() => {
      // すべてのタイマーをFast-forwardする
      jest.runAllTimers()
    })

    expect(result.current.toasts).toHaveLength(0)
  })
})
