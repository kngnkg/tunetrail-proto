export type Post = {
  id: string
  user_id: string
  user_name: string
  body: string
}

export type Pagination = {
  nextCursor: string
  previousCursor: string
  limit: number
}

export type Timeline = {
  posts: Post[]
  pagination: Pagination
}
