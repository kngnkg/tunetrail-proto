import fetcher from "@/lib/fetcher"

export const refreshAccessToken = async (apiRoot: string): Promise<void> => {
  return await fetcher(`${apiRoot}/auth/refresh`, {
    method: "POST",
    credentials: "include",
  })
}
