import { SignupForm } from "@/components/containers/SignupForm/SignupForm"

// ユーザー登録ページ
export default function SignupPage() {
  return (
    <div className="container mx-auto p-8">
      <h1 className="text-3xl mb-8">Signup Page</h1>
      <SignupForm />
    </div>
  )
}
