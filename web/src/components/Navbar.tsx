'use client'

import { HiDesktopComputer, HiMoon, HiSun } from 'react-icons/hi'
import { AnimatePresence, LayoutGroup, motion } from 'framer-motion'
import { useTheme } from 'next-themes'
import { usePathname } from 'next/navigation'
import { cn } from '~/utils/classNames'
import { siteConfig } from '~/config/site'
import useStore from '~/context/useStore'

export default function Navbar({ className }: { className?: string }) {
  const { theme, setTheme } = useTheme()
  const { titleInView } = useStore()
  const pathname = usePathname()

  return (
    <nav
      className={cn(
        'px-2 py-4 sm:px-4 lg:px-12 lg:py-6 flex justify-between items-center',
        className
      )}
    >
      {/* title */}
      <h1 className="text-xl md:text-3xl font-bold uppercase min-h-[22px] md:min-h-[29px]">
        <span className={cn(pathname === '/' && titleInView && 'hidden')}>{siteConfig.name}</span>
      </h1>
      {/* theme */}
      <LayoutGroup>
        <ul className={cn('flex items-center justify-center space-x-2 max-lg:order-1')}>
          {Object.entries({
            system: <HiDesktopComputer className={cn('h-4 w-auto')} />,
            dark: <HiMoon className={cn('h-4 w-auto')} />,
            light: <HiSun className={cn('h-4 w-auto')} />
          }).map(([key, value], i, self) => (
            <li
              key={key}
              className={cn('relative block cursor-pointer p-1.5')}
              onClick={() => setTheme(key)}
            >
              <AnimatePresence>
                {key === theme && (
                  <motion.div
                    layoutId="themeIdPointer"
                    initial={false}
                    className={cn(
                      'absolute inset-0 bg-neutral-800 dark:bg-neutral-500',
                      i === 0 && 'rounded-l-sm',
                      i === self.length - 1 && 'rounded-r-sm'
                    )}
                  />
                )}
              </AnimatePresence>
              <span
                className={cn('relative z-[1] block h-full w-full', {
                  'text-white': key === theme
                })}
              >
                {value}
              </span>
            </li>
          ))}
        </ul>
      </LayoutGroup>
    </nav>
  )
}
