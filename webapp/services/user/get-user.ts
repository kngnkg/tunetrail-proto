import ApiError from "@/types/error"
import { User, isUser } from "@/types/user"
import fetcher from "@/lib/fetcher"

const userNotFoundCode = 4201

export const getUser = async (
  apiRoot: string,
  userName: string
): Promise<User | null> => {
  try {
    const data = await fetcher(`${apiRoot}/users/${userName}`, {
      method: "GET",
      credentials: "include",
    })

    if (!isUser(data)) {
      throw new Error("Failed to fetch data from the API.")
    }

    return data
  } catch (e) {
    if (e instanceof ApiError) {
      switch (e.code) {
        case userNotFoundCode:
          return null
        default:
          throw e
      }
    }

    throw e
  }
}
