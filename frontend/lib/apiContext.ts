import { ApiContext } from "@/types"

export const apiContext: ApiContext = {
  apiRoot: process.env.NEXT_PUBLIC_API_ROOT || "",
}
