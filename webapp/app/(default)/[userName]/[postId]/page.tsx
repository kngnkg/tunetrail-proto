interface PostPageProps {
  params: { postId: string }
}

export default function PostPage({ params }: PostPageProps) {
  return (
    <div className="container mx-auto p-8">
      <h1 className="text-3xl mb-8">Post Page</h1>
      <p>{params.postId}</p>
    </div>
  )
}
