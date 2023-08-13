"use client"

import * as React from "react"

import { User, isUser } from "@/types/user"

export interface SigninData {
  userName: string
  email: string
  password: string
}
export const isSigninData = (arg: any): arg is SigninData => {
  return (
    (arg.userName !== undefined || arg.email !== undefined) &&
    arg.password !== undefined
  )
}

type SignedInUser = User

export type AuthContextType = {
  signedInUser: SignedInUser | null
  setSignedInUser: React.Dispatch<React.SetStateAction<SignedInUser | null>>
}

export const AuthContext = React.createContext<AuthContextType>(
  {} as AuthContextType
)

export interface AuthProviderProps {
  children: React.ReactNode
}

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [signedInUser, setSignedInUser] = React.useState<SignedInUser | null>(
    null
  )

  return (
    <AuthContext.Provider value={{ signedInUser, setSignedInUser }}>
      {children}
    </AuthContext.Provider>
  )
}
