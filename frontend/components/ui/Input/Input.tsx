"use client"

import * as React from "react"

import { mergeClasses } from "@/lib/utils"

export interface InputProps
  // labelをプレースホルダーとして表示するため、placeholderは使用しない
  extends Omit<React.InputHTMLAttributes<HTMLInputElement>, "placeholder"> {
  placeholderText?: string
}

export const Input: React.FC<InputProps> = ({
  className,
  type = "text",
  placeholderText,
  ...props
}) => {
  const [isFocused, setIsFocused] = React.useState(false)
  const [inputValue, setInputValue] = React.useState("")

  return (
    <div className="relative">
      <input
        type={type}
        className={mergeClasses(
          "bg-gray p-3 pt-4 pb-1 rounded-md hover:bg-gray-light focus:outline-none focus:bg-gray focus:ring-1 focus:ring-gray-light",
          className
        )}
        onFocus={() => setIsFocused(true)}
        onBlur={() => setIsFocused(false)}
        onChange={(e) => setInputValue(e.target.value)}
        {...props}
      />
      {/* labelをプレースホルダーとして使用する */}
      <label
        className={`text-gray-lightest absolute left-3 transition-all duration-100 pointer-events-none origin-left
            ${inputValue || isFocused ? "text-tiny top-0.5" : "text-sm top-3"}`}
      >
        {placeholderText}
      </label>
    </div>
  )
}

Input.displayName = "Input"
