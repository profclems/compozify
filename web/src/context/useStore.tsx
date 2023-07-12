'use client'

import { createContext, useState, useContext, useMemo, ReactNode } from 'react'
import { ThemeProvider as Theme } from 'next-themes'
import useMounted from '~/hooks/useMounted'

interface Store {
  titleInView: boolean
  setTitleInView: (value: boolean) => void
}

const StoreContext = createContext<Store>({
  titleInView: false,
  setTitleInView: () => {}
})

export function StoreProvider({ children }: { children: ReactNode }) {
  const mounted = useMounted()
  const [titleInView, setTitleInView] = useState(false)

  const memoizedValue = useMemo(
    () => ({
      titleInView,
      setTitleInView
    }),
    [titleInView]
  )

  if (!mounted) return null

  return (
    <StoreContext.Provider value={memoizedValue}>
      <Theme attribute="class" enableSystem={true} defaultTheme="system">
        {children}
      </Theme>
    </StoreContext.Provider>
  )
}

export default function useStore() {
  return useContext(StoreContext)
}
