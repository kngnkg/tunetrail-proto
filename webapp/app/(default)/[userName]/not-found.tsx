import Link from "next/link"

export default function NotFound() {
  return (
    <div className="container mx-auto p-8">
      <h1 className="text-3xl mb-8">Not Found</h1>
      <Link href="/home" className="text-primary">
        ホームに戻る
      </Link>
    </div>
  )
}
