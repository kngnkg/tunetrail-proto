"use client"

import * as React from "react"

import { useToast } from "@/hooks/toast/use-toast"
import {
  Toast,
  ToastProvider,
  ToastViewPort,
} from "@/components/ui/Toast/Toast"

export const Toaster: React.FC = () => {
  const { toasts } = useToast()

  return (
    <ToastProvider>
      {toasts &&
        toasts.map((toast) => {
          return (
            <Toast
              key={toast.id}
              intent={toast.intent}
              content={toast.description}
              open={toast.open}
            />
          )
        })}
      <ToastViewPort position="top" />
    </ToastProvider>
  )
}
