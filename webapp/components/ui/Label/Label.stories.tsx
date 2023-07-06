import { Meta, StoryObj } from "@storybook/react"

import { Label } from "./Label"

const meta: Meta<typeof Label> = {
  title: "UI/Label",
  component: Label,
  tags: ["autodocs"],
  argTypes: {},
}

export default meta
type Story = StoryObj<typeof Label>

export const Default: Story = {
  args: {
    htmlFor: "input",
    children: "Default Label",
  },
}
