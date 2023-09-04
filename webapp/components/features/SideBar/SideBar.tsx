"use client"

import * as React from "react"
import Link from "next/link"

import { mergeClasses } from "@/lib/utils"
import { useSignedInUser } from "@/hooks/auth/use-signedin-user"
import { Button } from "@/components/ui/Button/Button"

import { NewPostDialog } from "../NewPostDialog/NewPostDialog"

interface SideBarProps extends React.HTMLAttributes<HTMLDivElement> {}

export const SideBar: React.FC<SideBarProps> = ({ className, ...props }) => {
  const signedInUser = useSignedInUser()

  return (
    <div
      className={mergeClasses("flex flex-col space-y-8", className)}
      {...props}
    >
      <Link href="/home" className="text-2xl">
        TuneTrail
      </Link>
      <div className="flex flex-col space-y-4">
        <Link href="/home" className="text-2xl text-gray-lightest">
          Home
        </Link>
        <Link href="/notification" className="text-2xl text-gray-lightest">
          Notification
        </Link>
        <Link
          href={`/${signedInUser?.userName}`}
          className="text-2xl text-gray-lightest"
        >
          Profile
        </Link>
      </div>
      <NewPostDialog>
        <Button>Post</Button>
      </NewPostDialog>
    </div>
  )
}
