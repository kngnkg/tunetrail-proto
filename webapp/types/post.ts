import { User } from "@/types/user"

export type Post = {
  id: string
  parentId: string
  user: User
  body: string
  likesCount: number
  liked: boolean
  createdAt: Date
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
