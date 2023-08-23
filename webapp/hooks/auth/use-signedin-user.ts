"use client"

import * as React from "react"
import { AuthContext } from "@/providers/AuthProvider"

import { User } from "@/types/user"

export const useSetSignedInUser = (): React.Dispatch<
  React.SetStateAction<User | null>
> => React.useContext(AuthContext).setSignedInUser

export const useSignedInUser = (): User | null =>
  React.useContext(AuthContext).signedInUser
