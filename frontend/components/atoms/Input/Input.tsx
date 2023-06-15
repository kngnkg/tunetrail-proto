import { HTMLAttributes, forwardRef } from "react"

import { mergeClasses } from "@/lib/utils"

export interface InputProps extends HTMLAttributes<HTMLInputElement> {}

export const Input = forwardRef<HTMLInputElement, InputProps>(
  ({ className, ...props }, ref) => {
    return (
      <input
        className={mergeClasses(
          className,
          "bg-gray-700 p-2 rounded-md focus:outline-none focus:ring-1 focus:ring-primary"
        )}
        ref={ref}
        {...props}
      />
    )
  }
)

Input.displayName = "Input"
