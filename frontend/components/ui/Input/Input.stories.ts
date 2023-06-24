import { Meta, StoryObj } from "@storybook/react"

import { Input } from "./Input"

const meta: Meta<typeof Input> = {
  title: "UI/Input",
  component: Input,
  tags: ["autodocs"],
  argTypes: {},
}

export default meta
type Story = StoryObj<typeof Input>

export const Default: Story = {
  args: {
    placeholderText: "Type something...",
  },
}

export const WithBorder: Story = {
  args: {
    placeholderText: "Type something...",
    className: "border border-primary",
  },
}

export const Disabled: Story = {
  args: {
    placeholderText: "Type something...",
    disabled: true,
  },
}

export const Email: Story = {
  args: {
    placeholderText: "メールアドレス",
    type: "email",
  },
}

export const Password: Story = {
  args: {
    placeholderText: "パスワード",
    type: "password",
  },
}
