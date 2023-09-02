"use client"

import * as React from "react"
import * as DialogPrimitive from "@radix-ui/react-dialog"
import { Cross2Icon } from "@radix-ui/react-icons"

import { mergeClasses } from "@/lib/utils"

export const Dialog = DialogPrimitive.Root

export const DialogTrigger = DialogPrimitive.Trigger

export interface DialogContentProps
  extends React.ComponentPropsWithRef<typeof DialogPrimitive.Content> {
  className?: string
}

export const DialogContent = React.forwardRef<
  HTMLDivElement,
  DialogContentProps
>(({ className, children, ...props }, ref) => {
  return (
    <DialogPrimitive.Portal>
      <div className="fixed inset-0 z-50 flex items-start justify-center sm:items-center">
        <DialogPrimitive.Overlay
          ref={ref}
          className="fixed inset-0 z-50 bg-background/80 backdrop-blur-sm transition-all duration-100 data-[state=closed]:animate-out data-[state=closed]:fade-out data-[state=open]:fade-in"
        />
        <DialogPrimitive.Content
          ref={ref}
          className={mergeClasses(
            "fixed z-50 grid w-full gap-4 rounded-b-lg border bg-background p-6 shadow-lg animate-in data-[state=open]:fade-in-90 data-[state=open]:slide-in-from-bottom-10 sm:max-w-lg sm:rounded-lg sm:zoom-in-90 data-[state=open]:sm:slide-in-from-bottom-0",
            className
          )}
          {...props}
        >
          {/* 閉じるボタン */}
          <DialogPrimitive.Close asChild>
            <button className="IconButton" aria-label="Close">
              <Cross2Icon />
            </button>
          </DialogPrimitive.Close>
          {children}
        </DialogPrimitive.Content>
      </div>
    </DialogPrimitive.Portal>
  )
})
DialogContent.displayName = "DialogContent"
