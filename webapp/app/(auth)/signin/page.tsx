import Link from "next/link"

import { SigninForm } from "@/components/containers/SigninForm/SigninForm"

// サインインページ
export default function SigninPage() {
  return (
    <div className="container mx-auto p-8">
      <SigninForm />
      <div className="p-8">
        アカウントをお持ちでないですか？
        <Link href="/signup" className="text-primary">
          登録する
        </Link>
      </div>
    </div>
  )
}
