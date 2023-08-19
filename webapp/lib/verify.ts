import jwt, {
  JwtHeader,
  SigningKeyCallback,
  TokenExpiredError,
} from "jsonwebtoken"
import jwksClient from "jwks-rsa"

import { env } from "@/env.mjs"

var client = jwksClient({
  jwksUri: `https://cognito-idp.${env.TUNETRAIL_AWS_REGION}.amazonaws.com/${env.COGNITO_USER_POOL_ID}/.well-known/jwks.json`,
})

function getKey(header: JwtHeader, callback: SigningKeyCallback) {
  if (!header.kid) {
    throw new Error("not found kid!")
  }

  client.getSigningKey(header.kid, function (err, key) {
    if (err) {
      throw err
    }

    if (!key) {
      throw new Error("not found key!")
    }

    callback(null, key.getPublicKey())
  })
}

// function validateClaims(decoded: any): boolean {
//   // 有効期限の検証
//   if (decoded.exp && Date.now() >= decoded.exp * 1000) {
//     console.error("Token has expired.")
//     return false
//   }

//     // 発行者の検証
//     if (decoded.iss !== COGNITO_POOL_URL) {
//       console.error("Token issuer mismatch.")
//       return false
//     }

//   // TODO: その他の検証

//   return true
// }

export async function verifyToken(token: string): Promise<boolean> {
  return new Promise((resolve, reject) => {
    jwt.verify(token, getKey, (err, decoded) => {
      if (err) {
        if (err instanceof TokenExpiredError) {
          // アクセストークン切れの場合と処理を分ける場合、何かを書く
          resolve(false)
        } else {
          reject(err)
        }
      } else {
        // TODO: あなたのトークンの検証ロジックをここに追加します。
        // 例: return validateClaims(decoded)
        resolve(true)
      }
    })
  })
}
