import { ApiContext } from "@/types"

export const apiContext: ApiContext = {
  clientApiRoot: process.env.NEXT_PUBLIC_API_ROOT || "",
  serverApiRoot: process.env.API_ROOT || "",
}
