'use client'

// https://github.com/streamich/react-use/blob/master/src/useMedia.ts
import { useEffect, useState } from 'react'

/**
 * useMedia hook to detect media queries
 * @param {string} query - media query to evaluate
 * @param {boolean} defaultState - default state
 * @returns boolean
 * @example
 * const isWide = useMedia('(min-width: 480px)')
 * // isWide is true if screen width is >= 480px
 *
 */
export default function useMedia(query: string, defaultState: boolean = false): boolean {
  const [state, setState] = useState(defaultState)

  useEffect(() => {
    let mounted = true
    const mql = window.matchMedia(query)
    const onChange = () => {
      if (!mounted) return

      setState(mql.matches)
    }

    mql.addListener(onChange)
    setState(mql.matches)

    return () => {
      mounted = false
      mql.removeListener(onChange)
    }
  }, [query])

  return state
}
