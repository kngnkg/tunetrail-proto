import * as z from "zod"

import { MESSAGE } from "@/config/messages"

export const postSchema = z.object({
  body: z
    .string()
    .min(1, MESSAGE.VALIDATION.POST.BODY_TOO_SHORT)
    .max(1000, MESSAGE.VALIDATION.POST.BODY_TOO_LONG),
})
