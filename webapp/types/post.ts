import { User } from "@/types/user"

export type Post = {
  id: string
  user: User
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
