import { Meta, StoryObj } from "@storybook/react"

import { LoginForm } from "./LoginForm"

const meta: Meta<typeof LoginForm> = {
  title: "Containers/UserAuthForm",
  component: LoginForm,
  tags: ["autodocs"],
  argTypes: {},
}

export default meta
type Story = StoryObj<typeof LoginForm>

export const Default: Story = {
  args: {
    placeholder: "Type something...",
  },
}
