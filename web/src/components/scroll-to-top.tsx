'use client'

import { useEffect, useRef } from 'react'
import { usePathname } from 'next/navigation'
import useTop from '~/hooks/useTop'
import clsx from 'clsx'
import { HiArrowUp } from 'react-icons/hi'

export default function ScrollTop({
  className,
  disableOnRoutes,
  disableOnLayouts
}: {
  className?: string
  disableOnRoutes?: string[]
  disableOnLayouts?: string[]
}) {
  const scrollRef = useRef<HTMLButtonElement>(null)
  const pathname = usePathname()

  useEffect(() => {
    // Scroll event handler
    function scroll() {
      // Check if ref exists
      if (scrollRef.current) {
        // Check if scroll position is greater than 100px
        if (window.scrollY > 100) {
          scrollRef.current.classList.remove('translate-y-20')
        } else {
          scrollRef.current.classList.add('translate-y-20')
        }
      }
    }

    // Add event listener
    window.addEventListener('scroll', scroll)

    return () => {
      // Remove event listener
      window.removeEventListener('scroll', scroll)
    }
  }, [])

  return (
    <button
      ref={scrollRef}
      className={clsx(
        className,
        'flex h-10 w-10 translate-y-20 items-center justify-center rounded-sm bg-neutral-900 text-white transition-transform dark:bg-neutral-500 dark:text-white',
        disableOnRoutes && disableOnRoutes.map(route => route === pathname && 'hidden'),
        disableOnLayouts && disableOnLayouts.map(layout => pathname.startsWith(layout) && 'hidden')
      )}
      onClick={() => {
        window.scrollTo({
          top: 0,
          behavior: 'smooth'
        })
      }}
    >
      <HiArrowUp className="h-5 w-auto" />
    </button>
  )
}

export function ScrollToTopWithDocs({
  className,
  disableOnRoutes,
  disableOnLayouts
}: {
  className?: string
  disableOnRoutes?: string[]
  disableOnLayouts?: string[]
}) {
  const pathname = usePathname()
  const top = useTop()

  return (
    <button
      className={clsx(
        className,
        'group/link relative inline-flex items-center space-x-2 pb-1.5 pl-0.5 uppercase text-neutral-600 dark:text-neutral-400',
        disableOnRoutes && disableOnRoutes.map(route => route === pathname && 'hidden'),
        disableOnLayouts && disableOnLayouts.map(layout => pathname.startsWith(layout) && 'hidden'),
        top < 100 && 'hidden'
      )}
      onClick={() => {
        window.scrollTo({
          top: 0,
          behavior: 'smooth'
        })
      }}
    >
      <HiArrowUp className="h-3 w-auto" />
      <span>Scroll to top</span>
    </button>
  )
}
