'use client'

import Link from 'next/link'
import { usePathname } from 'next/navigation'
import MobileSideNav from '~/components/mobile-side-nav'
import { siteConfig } from '~/config/site'
import useStore from '~/context/useStore'
import { cn } from '~/utils/classNames'
import { AnimatePresence, LayoutGroup, motion } from 'framer-motion'
import { useTheme } from 'next-themes'
import { FaGithub } from 'react-icons/fa'
import { HiDesktopComputer, HiMoon, HiSun } from 'react-icons/hi'

export default function Navbar({ className }: { className?: string }) {
  const { theme, setTheme } = useTheme()
  const { titleInView } = useStore()
  const pathname = usePathname()

  return (
    <nav className={cn('flex items-center justify-between px-2 py-4 sm:px-4 lg:px-12 lg:py-6', className)}>
      <div className="flex items-center space-x-2">
        {/* mobile side nav */}
        <MobileSideNav />
        {/* title */}
        <Link
          href="/"
          className={cn(
            'min-h-[22px] text-xl font-bold uppercase md:min-h-[29px] md:text-3xl',
            pathname === '/' && titleInView && 'hidden'
          )}
        >
          <span className="">{siteConfig.name}</span>
        </Link>
      </div>
      {/* menu */}
      <div
        className={cn(
          'flex items-center space-x-3',
          pathname === '/' && titleInView && 'justify-between md:w-full md:flex-1'
        )}
      >
        <ul className={cn('flex items-center space-x-3 max-md:hidden')}>
          <Link href={siteConfig.links.github} target="_blank" rel="noopener noreferrer">
            <FaGithub className={cn('h-6 w-auto')} />
          </Link>
          <Link
            href="/docs/installation"
            className={cn('relative px-2 py-1', pathname.startsWith('/docs') && 'bg-zinc-800 text-white')}
          >
            Docs
          </Link>
        </ul>
        {/* theme */}
        <LayoutGroup>
          <ul className={cn('flex items-center justify-center space-x-2')}>
            {Object.entries({
              system: <HiDesktopComputer className={cn('h-4 w-auto')} />,
              dark: <HiMoon className={cn('h-4 w-auto')} />,
              light: <HiSun className={cn('h-4 w-auto')} />
            }).map(([key, value], i, self) => (
              <li key={key} className={cn('relative block cursor-pointer p-1.5')} onClick={() => setTheme(key)}>
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
      </div>
    </nav>
  )
}
