import * as React from "react"

import { mergeClasses } from "@/lib/utils"

export const Card = React.forwardRef<
  HTMLDivElement,
  React.HTMLAttributes<HTMLDivElement>
>(({ className, ...props }, ref) => {
  return (
    <div
      className={mergeClasses("border-b border-gray-light", className)}
      {...props}
    />
  )
})
Card.displayName = "Card"

export const CardHeader = React.forwardRef<
  HTMLDivElement,
  React.HTMLAttributes<HTMLDivElement>
>(({ className, ...props }, ref) => {
  return <div className={mergeClasses("", className)} {...props} />
})
CardHeader.displayName = "CardHeader"

export const CardContent = React.forwardRef<
  HTMLDivElement,
  React.HTMLAttributes<HTMLDivElement>
>(({ className, ...props }, ref) => {
  return <div className={mergeClasses("pb-2", className)} {...props} />
})
CardContent.displayName = "CardContent"

export const CardHooter = React.forwardRef<
  HTMLDivElement,
  React.HTMLAttributes<HTMLDivElement>
>(({ className, ...props }, ref) => {
  return <div className={mergeClasses("", className)} {...props} />
})
CardHooter.displayName = "CardHooter"
