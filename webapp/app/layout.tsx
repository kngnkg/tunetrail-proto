import { Toaster } from "@/components/features/Toaster/Toaster"

import "./globals.css"
import { Inter } from "next/font/google"
import { AuthProvider } from "@/providers/AuthProvider"
import { ToastProvider } from "@/providers/ToastProvider"

const inter = Inter({ subsets: ["latin"] })

export const metadata = {
  title: "TuneTrail",
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <AuthProvider>
          <ToastProvider>
            {children}
            <Toaster />
          </ToastProvider>
        </AuthProvider>
      </body>
    </html>
  )
}
