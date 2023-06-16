import { InputHTMLAttributes, forwardRef } from "react"

import { mergeClasses } from "@/lib/utils"

export interface InputProps extends InputHTMLAttributes<HTMLInputElement> {}

export const Input = forwardRef<HTMLInputElement, InputProps>(
  ({ className, type = "text", ...props }, ref) => {
    return (
      <input
        type={type}
        className={mergeClasses(
          "bg-gray-700 p-2 rounded-md focus:outline-none focus:ring-1 focus:ring-primary",
          className
        )}
        ref={ref}
        {...props}
      />
    )
  }
)

Input.displayName = "Input"
