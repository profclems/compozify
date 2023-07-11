import { useEffect, useState } from 'react'

/**
 * useMounted - checks if the component is mounted on the client / browser
 *
 * @returns {boolean} - true if the component is mounted
 */
export default function useMounted(): boolean {
  // keep track of whether the component is mounted
  const [mounted, setMounted] = useState(false)

  // set mounted to true on the client
  useEffect(() => setMounted(true), [])

  return mounted
}
