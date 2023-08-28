import { PostList } from "@/components/features/PostList/PostList"

// タイムライン
export default function Home() {
  return (
    <div className="container mx-auto p-8">
      <h1 className="text-3xl mb-8">Home Page (Timeline)</h1>
      <PostList />
    </div>
  )
}
