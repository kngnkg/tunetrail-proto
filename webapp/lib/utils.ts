import { twMerge } from "tailwind-merge"
import { ClassNameValue } from "tailwind-merge/dist/lib/tw-join"

// mergeClassesはTailwindCSSのクラス名をマージする関数
export function mergeClasses(...classLists: ClassNameValue[]): string {
  return twMerge(...classLists)
}
