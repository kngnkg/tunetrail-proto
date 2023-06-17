"use client"

import { HTMLAttributes, useState } from "react"
import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import * as z from "zod"

import { mergeClasses } from "@/lib/utils"
import { userAuthSchema } from "@/lib/validations/auth.schema"
import { Input } from "@/components/ui/Input/Input"

interface UserAuthFormProps extends HTMLAttributes<HTMLDivElement> {}

type FormData = z.infer<typeof userAuthSchema>

export default function UserAuthForm({
  className,
  ...props
}: UserAuthFormProps) {
  const { register, handleSubmit, formState } = useForm<FormData>({
    resolver: zodResolver(userAuthSchema),
  })
  const [isLoading, setIsLoading] = useState<boolean>(false)

  async function onSubmit(data: FormData) {
    setIsLoading(true)
    // TODO: ユーザー登録またはユーザー認証処理
    alert(JSON.stringify(data, null, 2))
    setIsLoading(false)
    // TODO: ユーザー登録またはユーザー認証後の処理
  }

  return (
    <div className={mergeClasses("grid gap-6", className)}>
      <form>
        <div>
          <Input
            type="email"
            placeholder="メールアドレスを入力してください。"
          />
        </div>
        <div>
          <Input type="password" placeholder="パスワードを入力してください。" />
        </div>
      </form>
    </div>
  )
}
