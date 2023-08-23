import Link from "next/link"

import { SignupForm } from "@/components/features/SignupForm/SignupForm"

// ユーザー登録ページ
export default function SignupPage() {
  return (
    <div className="container mx-auto p-8">
      <SignupForm />
      <div className="p-8">
        アカウントをお持ちですか？
        <Link href="/signin" className="text-primary">
          サインイン
        </Link>
      </div>
    </div>
  )
}
