'use client'

import { createContext, ReactNode, useCallback, useContext, useMemo, useState } from 'react'
import useMounted from '~/hooks/useMounted'
import { ThemeProvider as Theme } from 'next-themes'

interface Store {
  titleInView: boolean
  setTitleInView: (value: boolean) => void
  compose: () => void
  code?: string
  menu: boolean
  setMenu: (value: boolean) => void
}

const StoreContext = createContext<Store>({
  titleInView: false,
  setTitleInView: () => {},
  compose: () => {},
  code: undefined,
  menu: false,
  setMenu: () => {}
})

export function StoreProvider({ children }: { children: ReactNode }) {
  const mounted = useMounted()
  const [titleInView, setTitleInView] = useState(false)
  const [code, setCode] = useState<undefined | string>(undefined)
  const [menu, setMenu] = useState(false)

  const compose = useCallback(async () => {
    try {
      const response = await fetch('/api/parse', {
        mode: 'cors',
        method: 'GET',
        headers: { 'Content-Type': 'application/json' }
      })

      // throw error if response is not ok to trigger catch block
      if (!response.ok) throw new Error('Failed to generate response')

      // get response body and handle it here
    } catch (error) {
      // use switch statement to handle different error types
    }
  }, [])

  const memoizedValue = useMemo(
    () => ({
      titleInView,
      setTitleInView,
      compose,
      code,
      menu,
      setMenu
    }),
    [code, compose, titleInView, menu]
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
