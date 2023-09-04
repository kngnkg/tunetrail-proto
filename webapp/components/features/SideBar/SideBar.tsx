"use client"

import * as React from "react"
import Link from "next/link"
import { BellIcon, HomeIcon, PersonIcon } from "@radix-ui/react-icons"

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
        <Link href="/home" className="text-gray-lightest">
          <div className="flex items-center space-x-3">
            <HomeIcon className="h-6 w-6" />
            <p className="text-2xl">Home</p>
          </div>
        </Link>
        <Link href="/notification" className="text-gray-lightest">
          <div className="flex items-center space-x-3">
            <BellIcon className="h-6 w-6" />
            <p className="text-2xl">Notification</p>
          </div>
        </Link>
        <Link
          href={`/${signedInUser?.userName}`}
          className="text-gray-lightest"
        >
          <div className="flex items-center space-x-3">
            <PersonIcon className="h-6 w-6" />
            <p className="text-2xl">Profile</p>
          </div>
        </Link>
      </div>
      <NewPostDialog>
        <Button>Post</Button>
      </NewPostDialog>
    </div>
  )
}
