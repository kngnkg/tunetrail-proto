"use client"

import * as React from "react"
import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import * as z from "zod"

import { env } from "@/env.mjs"
import { Post } from "@/types/post"
import { MESSAGE } from "@/config/messages"
import { mergeClasses } from "@/lib/utils"
import { postSchema } from "@/lib/validations/post.schema"
import { useSignedInUser } from "@/hooks/auth/use-signedin-user"
import { useTimeline } from "@/hooks/post/use-timeline"
import { useToast } from "@/hooks/toast/use-toast"
import { Button } from "@/components/ui/Button/Button"
import {
  Dialog,
  DialogContent,
  DialogTrigger,
} from "@/components/ui/Dialog/Dialog"
import { TextArea } from "@/components/ui/TextArea/TextArea"
import { UserAvatar } from "@/components/features/UserAvatar/UserAvatar"

export interface NewPostDialogProps
  extends React.ComponentPropsWithoutRef<typeof Dialog> {
  className?: string
  parentPost?: Post
}

type FormData = z.infer<typeof postSchema>

export const NewPostDialog: React.FC<NewPostDialogProps> = ({
  className,
  children,
  ...props
}) => {
  const { register, handleSubmit, formState, reset } = useForm<FormData>({
    resolver: zodResolver(postSchema),
  })
  const { error, addPost, addReply } = useTimeline(env.NEXT_PUBLIC_API_ROOT)
  // const { addReply } = useReply(env.NEXT_PUBLIC_API_ROOT)
  const { showToast } = useToast()
  const [isLoading, setIsLoading] = React.useState<boolean>(false)
  const [open, setOpen] = React.useState<boolean>(false)
  const signedInUser = useSignedInUser()

  // アンマウント時にフォームをリセット
  React.useEffect(() => {
    if (!open) {
      reset()
    }
  }, [open, reset])

  if (!signedInUser) {
    return null
  }

  const onSubmit = async (data: FormData) => {
    setIsLoading(true)

    try {
      if (props.parentPost) {
        await addReply({ body: data.body, parentPost: props.parentPost })
      } else {
        await addPost(data)
      }

      showToast({
        intent: "success",
        description: MESSAGE.SUCCESS_POST,
      })
    } catch (e) {
      if (e instanceof Error) {
        showToast({
          intent: "error",
          description: e.message,
        })
      }

      showToast({
        intent: "error",
        description: MESSAGE.UNKNOWN_ERROR,
      })
    } finally {
      setOpen(false)
      setIsLoading(false)
    }
  }

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>{children}</DialogTrigger>
      <DialogContent className={mergeClasses("", className)} {...props}>
        <form
          onSubmit={handleSubmit(onSubmit)}
          className="flex flex-col gap-12"
        >
          <div className="">
            <TextArea
              id="body"
              placeholder="投稿内容"
              disabled={isLoading}
              {...register("body")}
            />
            {formState.errors.body && (
              <p className="px-1 text-xs text-red-500">
                {formState.errors.body.message}
              </p>
            )}
          </div>
          <div className="flex justify-between">
            <UserAvatar user={signedInUser} />
            <Button size="small" type="submit" className="h-10">
              Post
            </Button>
          </div>
        </form>
      </DialogContent>
    </Dialog>
  )
}
NewPostDialog.displayName = "NewPostDialog"
