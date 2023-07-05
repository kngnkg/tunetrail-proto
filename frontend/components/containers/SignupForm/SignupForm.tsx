"use client"

import * as React from "react"
import { isSignupData, signup } from "@/services/auth/signup"
import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import * as z from "zod"

import { apiContext } from "@/config/api-context"
import { MESSAGE } from "@/config/messages"
import { mergeClasses } from "@/lib/utils"
import { userAuthSchema } from "@/lib/validations/auth.schema"
import { useToast } from "@/hooks/toast/use-toast"
import { Button } from "@/components/ui/Button/Button"
import { Input } from "@/components/ui/Input/Input"

export interface SignupFormProps extends React.HTMLAttributes<HTMLDivElement> {}

type FormData = z.infer<typeof userAuthSchema>

// 登録ボタン押下時の処理
export const SignupForm: React.FC<SignupFormProps> = ({
  className,
  ...props
}) => {
  const { register, handleSubmit, formState } = useForm<FormData>({
    resolver: zodResolver(userAuthSchema),
  })
  const [isLoading, setIsLoading] = React.useState<boolean>(false)
  const { showToast } = useToast()

  const onSubmit = async (data: FormData) => {
    setIsLoading(true)

    if (!isSignupData(data)) {
      setIsLoading(false)
      return
    }
    const error = await signup(apiContext, data)

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
      description: MESSAGE.SUCCESS_SIGNUP,
    })
    setIsLoading(false)
  }

  return (
    <div className={mergeClasses("grid gap-6", className)} {...props}>
      <form onSubmit={handleSubmit(onSubmit)}>
        <div className="grid gap-2">
          <div className="grid gap-1">
            <Input
              id="name"
              type="name"
              placeholderText="アカウント名"
              disabled={isLoading}
              {...register("name")}
            />
            {formState.errors.name && (
              <p className="px-1 text-xs text-red-500">
                {formState.errors.name.message}
              </p>
            )}
          </div>
          <div className="grid gap-1">
            <Input
              id="userName"
              type="userName"
              placeholderText="ユーザー名"
              disabled={isLoading}
              {...register("userName")}
            />
            {formState.errors.userName && (
              <p className="px-1 text-xs text-red-500">
                {formState.errors.userName.message}
              </p>
            )}
          </div>
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
            <Input
              id="confirmPassword"
              type="password"
              placeholderText="パスワード(再度入力)"
              disabled={isLoading}
              {...register("confirmPassword")}
            />
            {formState.errors.confirmPassword && (
              <p className="px-1 text-xs text-red-500">
                {formState.errors.confirmPassword.message}
              </p>
            )}
          </div>
          <div className="grid gap-1">
            <Button type="submit" disabled={isLoading}>
              登録
            </Button>
          </div>
        </div>
      </form>
    </div>
  )
}
