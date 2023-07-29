'use client'

import { createContext, ReactNode, useCallback, useContext, useEffect, useMemo, useState } from 'react'
import { ThemeProvider as Theme } from 'next-themes'
import { stringify } from 'yaml'
import useMounted from '~/hooks/useMounted'
import { ErrorCause } from '~/types/nav'

interface Store {
  titleInView: boolean
  setTitleInView: (value: boolean) => void
  compose: (command: string) => Promise<void>
  code?: string
  menu: boolean
  setMenu: (value: boolean) => void
  error?: Err
}

type Err = {
  message: string
  code: number
}

const StoreContext = createContext<Store>({
  titleInView: false,
  setTitleInView: () => {},
  compose: async () => {},
  code: undefined,
  menu: false,
  setMenu: () => {}
})

export function StoreProvider({ children }: { children: ReactNode }) {
  const mounted = useMounted()
  const [titleInView, setTitleInView] = useState(false)
  const [code, setCode] = useState<undefined | string>(undefined)
  const [menu, setMenu] = useState(false)
  const [error, setError] = useState<undefined | Err>(undefined)

  useEffect(() => {
    const e = setTimeout(() => setError(undefined), 5000)

    return () => clearTimeout(e)
  }, [error])

  const compose = useCallback(async (command: string) => {
    try {
      const response = await fetch('/api/parse', {
        mode: 'cors',
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ command: command })
      })

      // throw error if response is not ok to trigger catch block
      if (!response.ok) throw new Error('Failed to generate response')

      type Res = undefined | { output?: string }
      // get response body and handle it here
      const body: Res = await response.json()
      const str = body && typeof body?.output === 'string' ? body.output.replace(/^\s*\|/, '') : undefined
      setCode(body && body.output ? stringify(str) : undefined)
    } catch (error) {
      let err: ErrorCause
      if (error instanceof Error) err = error as ErrorCause
      else err = new Error('Error creating catalog', { cause: { error } }) as ErrorCause
      // use switch statement to handle different error types
      switch (err.cause?.res?.status) {
        case 400: // bad request - invalid data
          setError({ message: 'Bad request', code: 400 })
          break

        case 404: // not found - resource not found
          setError({ message: 'Not found', code: 404 })
          break

        case 500: // internal server error - unknown error
          setError({ message: 'Internal server error', code: 500 })
          break

        default:
          setError({ message: 'Unknown error', code: 0 })
          break
      }
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
