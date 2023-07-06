import { SigninForm } from "@/components/containers/SigninForm/SigninForm"

// ログインページ
export default function LoginPage() {
  return (
    <div className="container mx-auto p-8">
      <h1 className="text-3xl mb-8">Login Page</h1>
      <SigninForm />
    </div>
  )
}
