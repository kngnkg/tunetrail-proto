"use client"

import * as React from "react"
import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import * as z from "zod"

import { env } from "@/env.mjs"
import { postSchema } from "@/lib/validations/post.schema"
import { usePosts } from "@/hooks/post/use-posts"
import { Button } from "@/components/ui/Button/Button"
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogTrigger,
} from "@/components/ui/Dialog/Dialog"
import { Input } from "@/components/ui/Input/Input"

export interface NewPostDialogProps
  extends React.ComponentPropsWithoutRef<typeof Dialog> {
  className?: string
}

type FormData = z.infer<typeof postSchema>

export const NewPostDialog: React.FC<NewPostDialogProps> = ({
  className,
  children,
  ...props
}) => {
  const { register, handleSubmit, formState } = useForm<FormData>({
    resolver: zodResolver(postSchema),
  })
  const { data, error, addPost } = usePosts(env.NEXT_PUBLIC_API_ROOT)
  const [isLoading, setIsLoading] = React.useState<boolean>(false)

  const onSubmit = async (data: FormData) => {
    setIsLoading(true)

    await addPost(data)

    setIsLoading(false)
  }

  return (
    <Dialog>
      <DialogTrigger asChild>{children}</DialogTrigger>
      <DialogContent {...props}>
        <h1>Post Dialog!</h1>
        <form onSubmit={handleSubmit(onSubmit)}>
          <Input
            id="body"
            type="textarea"
            placeholderText="投稿内容を入力する"
            disabled={isLoading}
            {...register("body")}
          />
          {formState.errors.body && (
            <p className="px-1 text-xs text-red-500">
              {formState.errors.body.message}
            </p>
          )}
          <DialogClose asChild>
            <Button>Post</Button>
          </DialogClose>
        </form>
      </DialogContent>
    </Dialog>
  )
}
NewPostDialog.displayName = "NewPostDialog"
