import { FetchError, isApiError } from "@/types/error"
import { User, isUser } from "@/types/user"
import fetcher from "@/lib/fetcher"

export const getUser = async (
  apiRoot: string,
  userName: string
): Promise<User | FetchError> => {
  const data = await fetcher(`${apiRoot}/user/${userName}`, {
    method: "GET",
    credentials: "include",
  })

  if (!isUser(data)) {
    throw new Error("Failed to fetch data from the API.")
  }

  return data
}
