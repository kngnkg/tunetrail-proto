import * as React from "react"

import { mergeClasses } from "@/lib/utils"

export interface ButtonProps
  extends React.ButtonHTMLAttributes<HTMLButtonElement> {}

export const Button: React.FC<ButtonProps> = ({ className, ...props }) => {
  return (
    <button
      className={mergeClasses(
        "bg-primary-transparent border border-primary text-primary rounded-md hover:bg-gray-light",
        className
      )}
      {...props}
    />
  )
}
Button.displayName = "Button"
