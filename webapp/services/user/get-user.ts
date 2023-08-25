import ApiError from "@/types/error"
import { Tokens } from "@/types/tokens"
import { User, isUser } from "@/types/user"
import fetcher from "@/lib/fetcher"

const userNotFoundCode = 4201

export const getUser = async (
  apiRoot: string,
  tokens: Tokens,
  userName: string
): Promise<User | null> => {
  try {
    const data = await fetcher(`${apiRoot}/users/${userName}`, {
      cache: "no-store",
      headers: {
        // トークンの制御を考える
        Cookie: `idToken=${tokens.idToken}; accessToken=${tokens.accessToken};`,
      },
    })

    if (!isUser(data)) {
      throw new Error("Failed to fetch data from the API.")
    }

    console.log(data)

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
