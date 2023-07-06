import { Meta, StoryObj } from "@storybook/react"

import { SigninForm } from "./SigninForm"

const meta: Meta<typeof SigninForm> = {
  title: "Containers/UserAuthForm",
  component: SigninForm,
  tags: ["autodocs"],
  argTypes: {},
}

export default meta
type Story = StoryObj<typeof SigninForm>

export const Default: Story = {
  args: {
    placeholder: "Type something...",
  },
}
