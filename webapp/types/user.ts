export type User = {
  id: string
  userName: string
  name: string
  iconUrl: string
  bio: string
  isFollowing: boolean
  isFollowed: boolean
  createdAt: string
  updatedAt: string
}
export const isUser = (arg: any): arg is User => {
  return (
    arg.id !== undefined &&
    arg.userName !== undefined &&
    arg.name !== undefined &&
    arg.iconUrl !== undefined &&
    arg.bio !== undefined &&
    arg.createdAt !== undefined &&
    arg.updatedAt !== undefined
  )
}
