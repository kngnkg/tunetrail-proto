"use client"

import * as React from "react"
import { useRouter, useSearchParams } from "next/navigation"
import {
  confirmSignup,
  isConfirmSignupData,
} from "@/services/auth/confirmSignup"
import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import * as z from "zod"

import { env } from "@/env.mjs"
import { mergeClasses } from "@/lib/utils"
import { confirmSignupSchema } from "@/lib/validations/auth.schema"
import { useToast } from "@/hooks/toast/use-toast"
import { Button } from "@/components/ui/Button/Button"
import { Input } from "@/components/ui/Input/Input"

export interface ConfirmSignupFormProps
  extends React.HTMLAttributes<HTMLDivElement> {}

type FormData = z.infer<typeof confirmSignupSchema>

export const ConfirmSignupForm: React.FC<ConfirmSignupFormProps> = ({
  className,
  ...props
}) => {
  const router = useRouter()
  const searchParams = useSearchParams()

  const { register, handleSubmit, formState } = useForm<FormData>({
    resolver: zodResolver(confirmSignupSchema),
  })
  const [isLoading, setIsLoading] = React.useState<boolean>(false)
  const { showToast } = useToast()

  // 登録ボタン押下時の処理
  const onSubmit = async (data: FormData) => {
    setIsLoading(true)

    // userNameをクエリパラメータから取得
    const userName = searchParams.get("userName")
    if (!userName || !isConfirmSignupData({ userName, ...data })) {
      console.log("userName: ", userName)
      setIsLoading(false)
      return
    }

    const error = await confirmSignup(env.NEXT_PUBLIC_API_ROOT, {
      userName: userName,
      code: data.code,
    })
    if (error) {
      showToast({
        intent: "error",
        description: error,
      })
      setIsLoading(false)
      return
    }

    // サインインページに遷移する
    router.push("/signin")

    setIsLoading(false)
  }

  return (
    <div className={mergeClasses("grid gap-6", className)} {...props}>
      <form onSubmit={handleSubmit(onSubmit)}>
        <div className="grid gap-2">
          <div className="grid gap-1">
            <Input
              id="code"
              type="code"
              placeholderText="認証コード"
              disabled={isLoading}
              {...register("code")}
            />
            {formState.errors.code && (
              <p className="px-1 text-xs text-red-500">
                {formState.errors.code.message}
              </p>
            )}
          </div>
          <div className="grid gap-1">
            <Button type="submit" disabled={isLoading}>
              コードを送信
            </Button>
          </div>
        </div>
      </form>
    </div>
  )
}
