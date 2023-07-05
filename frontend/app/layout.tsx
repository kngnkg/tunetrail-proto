import { Toaster } from "@/components/containers/Toaster/Toaster"

import "./globals.css"
import { Inter } from "next/font/google"
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
        <ToastProvider>
          {children}
          <Toaster />
        </ToastProvider>
      </body>
    </html>
  )
}
