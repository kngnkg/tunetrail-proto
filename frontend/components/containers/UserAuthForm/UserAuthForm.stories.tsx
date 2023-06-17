import { Meta, StoryObj } from "@storybook/react"

import UserAuthForm from "./UserAuthForm"

const meta: Meta<typeof UserAuthForm> = {
  title: "Containers/UserAuthForm",
  component: UserAuthForm,
  tags: ["autodocs"],
  argTypes: {},
}

export default meta
type Story = StoryObj<typeof UserAuthForm>

export const Default: Story = {
  args: {
    placeholder: "Type something...",
  },
}
