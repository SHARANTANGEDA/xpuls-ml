import './globals.css'
import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import {Providers} from "@/app/providers";
import {Author} from "next/dist/lib/metadata/types/metadata-types";

const inter = Inter({ subsets: ['latin'] })

export const metadata: Metadata = {
  title: 'Xpuls | Integrated MLOps & LLMOps Platform',
  description: 'End-to-end MLOps & LLMOps Platform for building, deploying and monitoring models in Production',
    authors: {
      url: "https://www.linkedin.com/in/sai-sharan-tangeda/",
      name: "Sai Sharan Tangeda",
    },
    keywords: ["MLOps", "LLMOps", "Opensource", "Platform", "SAAS", "PAAS", "Kubernetes", "ML", "ChatGPT", "Llama2",
        "Langchain", "Llama-index"]
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body className={inter.className}>
      <Providers>
          <main className="light text-foreground bg-background">

          {children}
          </main>
      </Providers>

      </body>
    </html>
  )
}
