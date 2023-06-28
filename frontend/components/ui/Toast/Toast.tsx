import * as React from "react"
import * as ToastPrimitive from "@radix-ui/react-toast"
import { VariantProps, cva } from "class-variance-authority"

import { mergeClasses } from "@/lib/utils"

export const toastViewPortVariants = cva(
  "flex max-h-screen w-full flex-col-reverse p-4 sm:bottom-0 sm:right-0 sm:top-auto sm:flex-col md:max-w-[420px]",
  {
    variants: {
      position: {
        top: "fixed top-0 z-[100]",
        inPlace: "",
      },
    },
    defaultVariants: {
      position: "top",
    },
  }
)

export interface ToastViewPortProps
  extends React.ComponentPropsWithoutRef<typeof ToastPrimitive.Viewport>,
    VariantProps<typeof toastViewPortVariants> {}

/**
 * Toastの表示位置を指定するためのコンポーネント
 * @param className - Toastの表示領域に適用するクラス名
 * @param position - Toastの表示位置を指定するためのvariant
 * @param props - その他のToastPrimitive.Viewportのprops
 */
export const ToastViewPort: React.FC<ToastViewPortProps> = ({
  className,
  position,
  ...props
}) => {
  return (
    <ToastPrimitive.Viewport
      className={mergeClasses(toastViewPortVariants({ position }), className)}
      {...props}
    />
  )
}
ToastViewPort.displayName = "ToastViewPort"

export const toastVariants = cva(
  "animate-slide-bottom grid bg-gray-dark text-white rounded-lg border-s-4 shadow-lg p-2 w-60",
  {
    variants: {
      intent: {
        default: "border-primary",
        success: "border-green-500",
        error: "border-red-500",
      },
    },
    defaultVariants: {
      intent: "default",
    },
  }
)

export interface ToastProps
  extends React.ComponentPropsWithoutRef<typeof ToastPrimitive.Root>,
    VariantProps<typeof toastVariants> {
  title?: string
  content?: string
}

/**
 * トースト
 * @param className - Toastに適用するクラス名
 * @param intent - Toastの意図を指定するためのvariant
 * @param title - タイトル
 * @param content - 内容
 * @param props - その他のToastPrimitive.Rootのprops
 */
export const Toast: React.FC<ToastProps> = ({
  className,
  intent,
  title,
  content,
  ...props
}) => {
  return (
    <ToastPrimitive.Root
      className={mergeClasses(toastVariants({ intent }), className)}
      {...props}
    >
      {title && (
        <ToastPrimitive.Title className="grid-cols-2 px-2 font-bold">
          {title}
        </ToastPrimitive.Title>
      )}
      <ToastPrimitive.Description className="grid-cols-1 text-xs px-2">
        {content}
      </ToastPrimitive.Description>
    </ToastPrimitive.Root>
  )
}
Toast.displayName = "Toast"
