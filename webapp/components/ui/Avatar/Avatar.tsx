import React from "react"
import * as AvatarPrimitive from "@radix-ui/react-avatar"

import { mergeClasses } from "@/lib/utils"

export interface AvatarProps
  extends React.ComponentPropsWithoutRef<typeof AvatarPrimitive.Root> {
  className?: string
}

export const Avatar: React.FC<AvatarProps> = ({ className, ...props }) => {
  return <AvatarPrimitive.Root className={className} {...props} />
}

export const AvatarImage: React.FC<
  React.ComponentPropsWithoutRef<typeof AvatarPrimitive.Image>
> = ({ className, ...props }) => {
  return (
    <AvatarPrimitive.Image
      className={mergeClasses(
        "aspect-square h-10 w-10 rounded-full overflow-hidden",
        className
      )}
      {...props}
    />
  )
}

export const AvatarFallback: React.FC<
  React.ComponentPropsWithoutRef<typeof AvatarPrimitive.Fallback>
> = ({ className, ...props }) => {
  return (
    <AvatarPrimitive.Fallback
      className={mergeClasses(
        "flex h-full w-full items-center justify-center rounded-full bg-muted",
        className
      )}
      {...props}
    />
  )
}
