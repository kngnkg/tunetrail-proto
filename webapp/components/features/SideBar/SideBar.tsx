import Link from "next/link"

import { mergeClasses } from "@/lib/utils"
import { Button } from "@/components/ui/Button/Button"

interface SideBarProps extends React.HTMLAttributes<HTMLDivElement> {
  signedInUserName: string
}

export const SideBar: React.FC<SideBarProps> = ({
  className,
  signedInUserName,
  ...props
}) => {
  return (
    <div
      className={mergeClasses("flex flex-col space-y-4", className)}
      {...props}
    >
      <Link href="/home" className="text-2xl text-gray-lightest">
        Home
      </Link>
      <Link href="/notification" className="text-2xl text-gray-lightest">
        Notification
      </Link>
      <Link
        href={`/${signedInUserName}`}
        className="text-2xl text-gray-lightest"
      >
        Profile
      </Link>
      <Button>Post</Button>
    </div>
  )
}
