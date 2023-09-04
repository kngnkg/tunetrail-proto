import React from "react"
import { AvatarProps } from "@radix-ui/react-avatar"

import { User } from "@/types/user"
import {
  Avatar,
  AvatarFallback,
  AvatarImage,
} from "@/components/ui/Avatar/Avatar"

interface UserAvatarProps extends AvatarProps {
  user: Pick<User, "name" | "iconUrl">
  className?: string
}

export const UserAvatar: React.FC<UserAvatarProps> = ({
  user,
  className,
  ...props
}) => {
  return (
    <Avatar className={className} {...props}>
      {user.iconUrl ? (
        <AvatarImage src={user.iconUrl} alt={user.name} />
      ) : (
        <AvatarFallback>{/* TODO: 考える */}</AvatarFallback>
      )}
    </Avatar>
  )
}
