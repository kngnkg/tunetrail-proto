import * as z from "zod"

import { MESSAGE } from "@/config/messages"

// パスワードのバリデーションルールを定義する
// 最低8文字、最大64文字
// 少なくとも1つの小文字、大文字、数字、記号
const passwordSchema = z
  .string()
  .min(8, MESSAGE.VALIDATION.USER.PASSWORD_TOO_SHORT)
  .max(64, MESSAGE.VALIDATION.USER.PASSWORD_TOO_LONG)
  .refine(
    (value) =>
      /[a-z]+/.test(value) &&
      /[A-Z]+/.test(value) &&
      /[0-9]+/.test(value) &&
      /[!@#\$%\^&\*\(\)-_=\+{}\[\]:;"'<>,\.\?/\\~\|]/.test(value),
    MESSAGE.VALIDATION.USER.PASSWORD_COMPLEXITY
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

export const confirmSignupSchema = z.object({
  code: z
    .string()
    .min(6, "確認コードは6文字です")
    .max(6, "確認コードは6文字です")
    .regex(/^\d+$/, "確認コードは数字のみです"),
})

export const signInSchema = z.object({
  email: z.string().email(MESSAGE.VALIDATION.USER.EMAIL),
  password: passwordSchema,
})
