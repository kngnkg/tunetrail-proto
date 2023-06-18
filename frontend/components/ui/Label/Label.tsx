import * as React from "react"
import * as LabelPrimitive from "@radix-ui/react-label"

import { mergeClasses } from "@/lib/utils"

export interface LabelProps
  extends React.ComponentProps<typeof LabelPrimitive.Root> {
  className?: string
}

const Label: React.FC<LabelProps> = ({ className, ...props }) => (
  <LabelPrimitive.Root
    className={mergeClasses(
      "text-sm font-medium text-black dark:text-white",
      className
    )}
    {...props}
  />
)
Label.displayName = LabelPrimitive.Root.displayName

export { Label }
