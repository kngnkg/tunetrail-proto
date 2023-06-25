import { Meta, StoryObj } from "@storybook/react"

import { Button } from "./Button"

const meta: Meta<typeof Button> = {
  title: "UI/Button",
  component: Button,
  tags: ["autodocs"],
  argTypes: {},
}

export default meta
type Story = StoryObj<typeof Button>

export const Default: Story = {
  args: {
    children: "ボタン",
  },
}

export const Small: Story = {
  args: {
    children: "Small",
    size: "small",
  },
}

export const Large: Story = {
  args: {
    children: "Small",
    size: "large",
  },
}

export const Long: Story = {
  args: {
    children: "Long",
    size: "medium",
    className: "w-96",
  },
}

export const Disabled: Story = {
  args: {
    children: "Disabled",
    disabled: true,
  },
}
