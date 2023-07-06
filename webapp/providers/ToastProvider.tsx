"use client"

import * as React from "react"

import { ToastProps } from "@/components/ui/Toast/Toast"

export type Toast = ToastProps & {
  intent: "default" | "success" | "error"
  description: string
  open: boolean
  id: string
}

export type ToastContextType = {
  toasts: Toast[]
  setToasts: React.Dispatch<React.SetStateAction<Toast[]>>
}

export const ToastContext = React.createContext<ToastContextType>(
  {} as ToastContextType
)

export interface ToastProviderProps {
  children: React.ReactNode
}

export const ToastProvider: React.FC<ToastProviderProps> = ({ children }) => {
  const [toasts, setToasts] = React.useState<Toast[]>([])

  return (
    <ToastContext.Provider value={{ toasts, setToasts }}>
      {children}
    </ToastContext.Provider>
  )
}
