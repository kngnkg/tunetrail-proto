import * as React from "react"
import { Toast, ToastContext } from "@/providers/ToastProvider"

const TIMEOUT = 5000

interface UseToastProps {
  intent: "default" | "success" | "error"
  description: string
}

// トーストに一意のIDを付与するためのカウンター
let count = 0
const generateToastId = (): string => {
  count++
  return count.toString()
}

export const useToast = () => {
  const { toasts, setToasts } = React.useContext(ToastContext)

  // showToastはトーストを表示する関数
  const showToast = (value: UseToastProps) => {
    const toast: Toast = {
      ...value,
      open: true,
      id: generateToastId(),
    }
    const toastId = toast.id

    setToasts((toasts) => [...toasts, toast])

    // TIMEOUTの時間が経過したらトーストを非表示にする
    setTimeout(() => {
      setToasts((toasts) => toasts.filter((toast) => toast.id !== toastId))
    }, TIMEOUT)
  }

  return { toasts, showToast }
}
