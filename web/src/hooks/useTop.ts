'use client'

import { useEffect, useState } from 'react'

/**
 * useTop - Returns the current scroll position of the window
 * @returns {number} - The current scroll position of the window
 * @example
 * const top = useTop()
 * console.log(top)
 * // => 0
 * // => 100
 */
export default function useTop(): number {
  const [top, setTop] = useState(0)

  useEffect(() => {
    function onScroll() {
      setTop(window.scrollY)
    }

    window.addEventListener('scroll', onScroll)

    return () => {
      window.removeEventListener('scroll', onScroll)
    }
  }, [])

  return top
}
