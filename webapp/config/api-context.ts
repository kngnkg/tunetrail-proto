import { ApiContext } from "@/types/api-context"

export const apiContext: ApiContext = {
  apiRoot: process.env.NEXT_PUBLIC_API_ROOT || "",
}
