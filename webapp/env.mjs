import { createEnv } from "@t3-oss/env-nextjs"
import { z } from "zod"

export const env = createEnv({
  //   server: {
  //     COGNITO_CLIENT_ID: z.string(),
  //     COGNITO_CLIENT_SECRET: z.string(),
  //     COGNITO_ISSUER: z.string().url(),
  //   },
  client: {
    NEXT_PUBLIC_API_ROOT: z.string().url(),
  },
  runtimeEnv: {
    NEXT_PUBLIC_API_ROOT: process.env.NEXT_PUBLIC_API_ROOT,
    // COGNITO_CLIENT_ID: process.env.COGNITO_CLIENT_ID,
    // COGNITO_CLIENT_SECRET: process.env.COGNITO_CLIENT_SECRET,
    // COGNITO_ISSUER: process.env.COGNITO_ISSUER,
  },
})
