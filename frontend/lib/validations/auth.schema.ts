import * as z from "zod"

// パスワードのバリデーションルールを定義する
// 最低8文字、最大20文字
// 少なくとも1つの小文字、大文字、数字
// 英数字のみ
const passwordSchema = z
  .string()
  .min(8, "パスワードは8文字以上である必要があります")
  .max(20, "パスワードは20文字以下である必要があります")
  .regex(
    /[a-z]+/,
    "少なくとも1つの小文字、大文字、数字が含まれている必要があります"
  )
  .regex(
    /[A-Z]+/,
    "少なくとも1つの小文字、大文字、数字が含まれている必要があります"
  )
  .regex(
    /[0-9]+/,
    "少なくとも1つの小文字、大文字、数字が含まれている必要があります"
  )
  .regex(/^[a-zA-Z0-9]+$/, "パスワードは英数字のみである必要があります")

export const userAuthSchema = z.object({
  email: z.string().email("メールアドレスの形式が正しくありません"),
  password: passwordSchema,
})
