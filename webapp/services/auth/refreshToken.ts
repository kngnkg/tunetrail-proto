import { clientFetcher } from "@/lib/fetcher"

export const refreshToken = async (apiRoot: string): Promise<void> => {
  return await clientFetcher(`${apiRoot}/refresh`, {
    method: "POST",
  })
}
