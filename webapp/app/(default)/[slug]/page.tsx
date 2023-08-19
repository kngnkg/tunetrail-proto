// ユーザーページ
export default function UserPage({ params }: { params: { slug: string } }) {
  return (
    <div className="container mx-auto p-8">
      <h1 className="text-3xl mb-8">User Page: {params.slug}</h1>
    </div>
  )
}
