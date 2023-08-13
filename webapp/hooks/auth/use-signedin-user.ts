"use client"

import * as React from "react"
import { AuthContext, AuthContextType } from "@/providers/AuthProvider"

export const useSignedInUser = (): AuthContextType => {
  const { signedInUser, setSignedInUser } = React.useContext(AuthContext)

  return { signedInUser, setSignedInUser }
}
