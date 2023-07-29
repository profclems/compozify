import '~/styles/globals.css'

import { Metadata } from 'next'
import LocalFont from 'next/font/local'
import Navbar from '~/components/Navbar'
import { siteConfig } from '~/config/site'
import { StoreProvider } from '~/context/useStore'
import { cn } from '~/utils/classNames'

export const metadata: Metadata = {
  title: {
    default: siteConfig.name,
    template: `%s - ${siteConfig.name}`
  },
  description: siteConfig.description.long,
  keywords: ['docker', 'compose.yml', 'devops'],
  authors: [
    {
      name: 'profclems',
      url: 'https://github.com/profclems'
    }
  ],
  themeColor: [
    { media: '(prefers-color-scheme: light)', color: 'white' },
    { media: '(prefers-color-scheme: dark)', color: 'black' }
  ],
  icons: {
    icon: '/favicon.ico',
    shortcut: '/icons/icon.png',
    apple: '/icons/icon.png'
  },
  manifest: `/manifest.json`
}

const Satoshi = LocalFont({
  src: [{ path: './Satoshi-Variable.woff2', style: 'normal' }],
  variable: '--font-satoshi'
})

const Lobster = LocalFont({
  src: [{ path: './lobster.ttf', style: 'normal' }],
  variable: '--font-lobster'
})

const JetbrainsMono = LocalFont({
  src: [
    { path: './jetbrainsmono.ttf', style: 'normal' },
    { path: './jetbrainsmono-italic.ttf', style: 'italic' }
  ],
  variable: '--font-mono'
})

export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <html
      className={cn(
        'bg-white text-zinc-900 dark:bg-zinc-800 dark:text-white',
        Satoshi.variable,
        Lobster.variable,
        JetbrainsMono.variable
      )}
    >
      <head />
      <body className="bg-white text-zinc-900 dark:bg-zinc-800 dark:text-white">
        <StoreProvider>
          <main className="bg-white text-zinc-900 min-h-screen dark:bg-zinc-800 dark:text-white">
            <Navbar className={cn('sticky inset-x-0 top-0 z-[4] bg-white/90 dark:bg-zinc-800/90')} />
            <div className="relative w-full">{children}</div>
          </main>
        </StoreProvider>
      </body>
    </html>
  )
}
