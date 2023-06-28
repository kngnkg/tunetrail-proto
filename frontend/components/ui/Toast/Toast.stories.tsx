import * as React from "react"
import { ToastProvider } from "@/providers/ToastProvider"
import { Meta, StoryObj } from "@storybook/react"
import { set } from "zod"

import { Toast, ToastProps, ToastViewPort } from "./Toast"

const meta: Meta<typeof Toast> = {
  title: "UI/Toast",
  component: Toast,
  tags: ["autodocs"],
  argTypes: {},
  decorators: [
    (Story) => (
      <ToastProvider>
        <Story />
      </ToastProvider>
    ),
  ],
}

export default meta
type Story = StoryObj<typeof Toast>

export const Default: Story = (args: ToastProps) => {
  const [open, setOpen] = React.useState(false)
  const timerRef = React.useRef<NodeJS.Timeout | null>(null)

  // コンポーネントがアンマウントされたときにタイマーをクリア
  React.useEffect(() => {
    return () => {
      if (timerRef.current) {
        clearTimeout(timerRef.current)
      }
    }
  }, [])

  React.useEffect(() => {
    if (open) {
      // トーストが表示されているときに5秒後に非表示にするタイマー
      timerRef.current = setTimeout(() => setOpen(false), 5000)
    }
  }, [open])

  return (
    <>
      <button
        className="bg-gray rounded-md text-xs p-2"
        onClick={() => {
          setOpen(!open)
        }}
      >
        Show Toast
      </button>
      <Toast open={open} onOpenChange={setOpen} {...args} />
      <ToastViewPort />
    </>
  )
}

Default.args = {
  content: "Content",
}

export const WithTitle: Story = (args: ToastProps) => (
  <>
    <Toast {...args} />
    <ToastViewPort position="inPlace" />
  </>
)
WithTitle.args = {
  open: true,
  title: "Title",
  content: "Content",
}

export const Success: Story = (args: ToastProps) => (
  <>
    <Toast {...args} />
    <ToastViewPort position="inPlace" />
  </>
)
Success.args = {
  open: true,
  content: "登録に成功しました!",
  intent: "success",
}

export const Error: Story = (args: ToastProps) => (
  <>
    <Toast {...args} />
    <ToastViewPort position="inPlace" />
  </>
)
Error.args = {
  open: true,
  content: "登録に失敗しました",
  intent: "error",
}
