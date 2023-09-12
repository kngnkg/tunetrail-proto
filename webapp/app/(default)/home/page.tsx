import Link from "next/link"

import { PostList } from "@/components/features/PostList/PostList"
import { TimeLine } from "@/components/features/TimeLine/TimeLine"

// タイムライン
export default function Home() {
  return (
    <div className="container mx-auto p-8">
      {/* <h1 className="text-3xl mb-8">Home</h1> */}
      <TimeLine />
    </div>
  )
}
