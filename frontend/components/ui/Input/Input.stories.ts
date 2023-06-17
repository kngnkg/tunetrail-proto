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
    placeholder: "Type something...",
  },
}

export const WithBorder: Story = {
  args: {
    placeholder: "Type something...",
    className: "border border-primary",
  },
}

export const Disabled: Story = {
  args: {
    placeholder: "Type something...",
    disabled: true,
  },
}

export const Email: Story = {
  args: {
    placeholder: "Type your email...",
    type: "email",
  },
}

export const Password: Story = {
  args: {
    placeholder: "Type your password...",
    type: "password",
  },
}
