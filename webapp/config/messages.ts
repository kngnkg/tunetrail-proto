export const MESSAGE = {
  VALIDATION: {
    USER: {
      USERNAME: "ユーザー名は5文字以上である必要があります",
      NAME: "アカウント名は1文字以上である必要があります",
      EMAIL: "メールアドレスの形式が正しくありません",
      PASSWORD_COMPLEXITY:
        "パスワードは少なくとも1つの小文字、大文字、数字、記号が含まれている必要があります",
      PASSWORD_TOO_SHORT: "パスワードは8文字以上である必要があります",
      PASSWORD_TOO_LONG: "パスワードが長すぎます",
      PASSWORD_CONFIRM: "パスワードと確認用パスワードが一致しません",
    },
    POST: {
      BODY_TOO_SHORT: "ポストは1文字以上である必要があります",
      BODY_TOO_LONG: "ポストは1000文字以内である必要があります",
    },
  },
  SUCCESS_SIGNUP: "登録に成功しました!",
  SUCCESS_LOGIN: "ログインしました!",
  UNKNOWN_ERROR: "不明なエラーが発生しました",
}
