import { ConfirmSignupForm } from "@/components/containers/ConfirmSignupForm/ConfirmSignupForm"

// メールアドレス認証コード入力画面
export default function ConfirmPage() {
  return (
    <div className="container mx-auto p-8">
      <ConfirmSignupForm />
    </div>
  )
}
