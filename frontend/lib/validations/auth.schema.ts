import * as z from "zod"

import { MESSAGE } from "@/config/messages"

// パスワードのバリデーションルールを定義する
// 最低8文字、最大20文字
// 少なくとも1つの小文字、大文字、数字
// 英数字のみ
const passwordSchema = z
  .string()
  .min(8)
  .max(20)
  .refine(
    (value) =>
      /[a-z]+/.test(value) &&
      /[A-Z]+/.test(value) &&
      /[0-9]+/.test(value) &&
      /^[a-zA-Z0-9]+$/.test(value),
    MESSAGE.VALIDATION.USER.PASSWORD
  )

const passwordConfirmationSchema = z
  .object({
    password: passwordSchema,
    confirmPassword: z.string(),
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: "パスワードと確認用パスワードが一致しません",
  })

export const userAuthSchema = z
  .object({
    userName: z.string().min(5, MESSAGE.VALIDATION.USER.USERNAME),
    name: z.string().min(1, MESSAGE.VALIDATION.USER.NAME),
    email: z.string().email(MESSAGE.VALIDATION.USER.EMAIL),
    password: passwordSchema,
    confirmPassword: z.string(),
  })
  .refine((data) => passwordConfirmationSchema.safeParse(data).success, {
    message: "パスワードと確認用パスワードが一致しません",
  })
