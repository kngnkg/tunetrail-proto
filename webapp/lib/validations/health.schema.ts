import { z } from "zod"

const healthStatusEnum = z.enum(["green", "orange", "red"])

export const HealthSchema = z.object({
  health: healthStatusEnum,
  database: healthStatusEnum,
})

export type Health = z.infer<typeof HealthSchema>
