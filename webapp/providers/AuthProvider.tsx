"use client"

import * as React from "react"

import { User } from "@/types/user"

export type AuthContextType = {
  signedInUser: User | null
  setSignedInUser: React.Dispatch<React.SetStateAction<User | null>>
}

export const AuthContext = React.createContext<AuthContextType>(
  {} as AuthContextType
)

export interface AuthProviderProps {
  children: React.ReactNode
}

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [signedInUser, setSignedInUser] = React.useState<User | null>(null)

  return (
    <AuthContext.Provider value={{ signedInUser, setSignedInUser }}>
      {children}
    </AuthContext.Provider>
  )
}
