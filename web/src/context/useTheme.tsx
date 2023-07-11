'use client'

import { ThemeProvider as Theme } from 'next-themes'
import useMounted from '~/hooks/useMounted'

export default function ThemeProvider({ children }: { children: React.ReactNode }) {
  const mounted = useMounted()

  if (!mounted) return null

  return (
    <Theme attribute="class" enableSystem={true} defaultTheme="system">
      {children}
    </Theme>
  )
}
