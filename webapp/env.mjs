import { createEnv } from "@t3-oss/env-nextjs"
import { z } from "zod"

export const env = createEnv({
  server: {
    TUNETRAIL_AWS_REGION: z.string(),
    COGNITO_USER_POOL_ID: z.string(),
    ALLOWED_DOMAIN: z.string(),
    API_ROOT: z.string().url(),
  },
  client: {
    NEXT_PUBLIC_API_ROOT: z.string().url(),
    NEXT_PUBLIC_AUTH_API_ROOT: z.string().url(),
  },
  runtimeEnv: {
    NEXT_PUBLIC_API_ROOT: process.env.NEXT_PUBLIC_API_ROOT,
    NEXT_PUBLIC_AUTH_API_ROOT: process.env.NEXT_PUBLIC_AUTH_API_ROOT,
    TUNETRAIL_AWS_REGION: process.env.TUNETRAIL_AWS_REGION,
    COGNITO_USER_POOL_ID: process.env.COGNITO_USER_POOL_ID,
    ALLOWED_DOMAIN: process.env.ALLOWED_DOMAIN,
    API_ROOT: process.env.API_ROOT,
  },
})
