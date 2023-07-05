"use client"

import * as React from "react"
import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import * as z from "zod"

import { MESSAGE } from "@/config/messages"
import { mergeClasses } from "@/lib/utils"
import { userAuthSchema } from "@/lib/validations/auth.schema"
import { useToast } from "@/hooks/toast/use-toast"
import { Button } from "@/components/ui/Button/Button"
import { Input } from "@/components/ui/Input/Input"

interface UserAuthFormProps extends React.HTMLAttributes<HTMLDivElement> {}

type FormData = z.infer<typeof userAuthSchema>

export default function UserAuthForm({
  className,
  ...props
}: UserAuthFormProps) {
  const { register, handleSubmit, formState } = useForm<FormData>({
    resolver: zodResolver(userAuthSchema),
  })
  const [isLoading, setIsLoading] = React.useState<boolean>(false)
  const { showToast } = useToast()

  const onSubmit = async (data: FormData) => {
    setIsLoading(true)

    // TODO: ユーザー登録またはユーザー認証処理
    await new Promise((resolve) => setTimeout(resolve, 1000))
    const signinResult = confirm(`Signin with ${JSON.stringify(data)}`)

    // TODO: ユーザー登録またはユーザー認証後の処理
    if (!signinResult) {
      showToast({
        intent: "error",
        description: MESSAGE.LOGIN.FAILED,
      })
      setIsLoading(false)
      return
    }

    showToast({
      intent: "success",
      description: MESSAGE.LOGIN.SUCCESS,
    })
    setIsLoading(false)
  }

  return (
    <div className={mergeClasses("grid gap-6", className)} {...props}>
      <form onSubmit={handleSubmit(onSubmit)}>
        <div className="grid gap-2">
          <div className="grid gap-1">
            <Input
              id="email"
              type="email"
              placeholderText="メールアドレス"
              disabled={isLoading}
              {...register("email")}
            />
            {formState.errors.email && (
              <p className="px-1 text-xs text-red-500">
                {formState.errors.email.message}
              </p>
            )}
          </div>
          <div className="grid gap-1">
            <Input
              id="password"
              type="password"
              placeholderText="パスワード"
              disabled={isLoading}
              {...register("password")}
            />
            {formState.errors.password && (
              <p className="px-1 text-xs text-red-500">
                {formState.errors.password.message}
              </p>
            )}
          </div>
          <div className="grid gap-1">
            <Button type="submit" disabled={isLoading}>
              ログイン
            </Button>
          </div>
        </div>
      </form>
    </div>
  )
}
