"use client"

import * as React from "react"
import { useRouter } from "next/navigation"
import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import * as z from "zod"

import { env } from "@/env.mjs"
import { MESSAGE } from "@/config/messages"
import { mergeClasses } from "@/lib/utils"
import { signInSchema } from "@/lib/validations/auth.schema"
import { useSignin } from "@/hooks/auth/use-signin"
import { useToast } from "@/hooks/toast/use-toast"
import { Button } from "@/components/ui/Button/Button"
import { Input } from "@/components/ui/Input/Input"

interface SigninFormProps extends React.HTMLAttributes<HTMLDivElement> {}

type FormData = z.infer<typeof signInSchema>

export const SigninForm: React.FC<SigninFormProps> = ({
  className,
  ...props
}) => {
  const router = useRouter()
  const { register, handleSubmit, formState } = useForm<FormData>({
    resolver: zodResolver(signInSchema),
  })

  const [isLoading, setIsLoading] = React.useState<boolean>(false)
  const { error, signin } = useSignin()
  const { showToast } = useToast()

  const onSubmit = async (data: FormData) => {
    setIsLoading(true)

    await signin(env.NEXT_PUBLIC_API_ROOT, {
      userName: "", // TODO: ユーザー名を入力できるようにする
      ...data,
    })
    if (error) {
      showToast({
        intent: "error",
        description: error,
      })
      setIsLoading(false)
      return
    }

    showToast({
      intent: "success",
      description: MESSAGE.SUCCESS_LOGIN,
    })

    // サインイン後はホーム画面に遷移
    router.push("/")

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
              サインイン
            </Button>
          </div>
        </div>
      </form>
    </div>
  )
}
