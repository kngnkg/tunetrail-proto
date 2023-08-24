"use client"

import * as React from "react"
import { useRouter } from "next/navigation"

import { env } from "@/env.mjs"
import { useSignedInUser } from "@/hooks/auth/use-signedin-user"
import { useToast } from "@/hooks/toast/use-toast"
import { useFollow } from "@/hooks/user/use-follow"
import { Button, ButtonProps } from "@/components/ui/Button/Button"

export interface FollowButtonProps extends ButtonProps {
  userName: string
  isFollowing: boolean
}

export const FollowButton: React.FC<FollowButtonProps> = ({
  className,
  userName,
  isFollowing: initialIsFollowing,
  ...props
}) => {
  const router = useRouter()
  const [isLoading, setIsLoading] = React.useState(false)

  const { showToast } = useToast()
  const { isFollowing, error, follow, unfollow } = useFollow({
    apiRoot: env.NEXT_PUBLIC_API_ROOT,
    isFollowing: initialIsFollowing,
  })
  const signedInUser = useSignedInUser()

  React.useEffect(() => {
    console.log("signedInUser")
    console.log(isFollowing)

    if (signedInUser) {
      return
    }

    router.push("/signin")
  }, [signedInUser, router])

  const onClickToggleFollowing = async () => {
    if (!signedInUser) {
      return
    }

    setIsLoading(true)

    if (isFollowing) {
      await unfollow(signedInUser.userName, userName)
    } else {
      await follow(signedInUser.userName, userName)
    }

    if (error) {
      showToast({
        intent: "error",
        description: error,
      })
    }

    setIsLoading(false)
  }

  return (
    <Button
      onClick={onClickToggleFollowing}
      intent="primary"
      size="medium"
      className={className}
      disabled={isLoading}
      {...props}
    >
      {isFollowing ? "フォロー中" : "フォローする"}
    </Button>
  )
}
