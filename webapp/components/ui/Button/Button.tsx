import * as React from "react"
import { VariantProps, cva } from "class-variance-authority"

import { mergeClasses } from "@/lib/utils"

export const buttonVariants = cva("text-sm disabled:pointer-events-none", {
  variants: {
    intent: {
      primary:
        "bg-primary-transparent border border-primary text-primary rounded-md hover:bg-gray-light",
    },
    size: {
      small: "w-20 text-xs",
      medium: "w-40 h-10",
      large: "w-60 h-12",
    },
  },
  defaultVariants: {
    intent: "primary",
    size: "medium",
  },
})

export interface ButtonProps
  extends React.ButtonHTMLAttributes<HTMLButtonElement>,
    VariantProps<typeof buttonVariants> {}

export const Button = React.forwardRef<HTMLButtonElement, ButtonProps>(
  ({ className, intent, size, ...props }, ref) => {
    return (
      <button
        className={mergeClasses(buttonVariants({ intent, size }), className)}
        {...props}
        ref={ref}
      />
    )
  }
)
Button.displayName = "Button"
