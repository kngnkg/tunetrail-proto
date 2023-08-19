import fetcher from "@/lib/fetcher"

export const refreshToken = async (apiRoot: string): Promise<void> => {
  return await fetcher(`${apiRoot}/refresh`, {
    method: "POST",
    credentials: "include",
  })
}
