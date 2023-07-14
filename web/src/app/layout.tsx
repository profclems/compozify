import '~/styles/globals.css'
import { Metadata } from 'next'
import LocalFont from 'next/font/local'
import Image from 'next/image'
import { cn } from '~/utils/classNames'
import { siteConfig } from '~/config/site'
import { StoreProvider } from '~/context/useStore'
import Navbar from '~/components/Navbar'

export const metadata: Metadata = {
  title: {
    default: siteConfig.name,
    template: `%s - ${siteConfig.name}`
  },
  description: siteConfig.description,
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
          <main className="bg-white text-zinc-900 dark:bg-zinc-800 dark:text-white">
            <div className={cn('relative')}>
              <div className="fixed inset-0">
                <Image
                  src="/images/beams-2.png"
                  alt="Background parttern"
                  loading="eager"
                  fill
                  style={{
                    position: 'absolute',
                    inset: 0,
                    width: '100%',
                    height: '100%',
                    transform: 'rotate(45deg)'
                  }}
                />
              </div>
              <Navbar className={cn('fixed z-[4] top-0 inset-x-0')} />
              <div className="relative inset-0 z-[3] min-h-screen w-full">{children}</div>
            </div>
          </main>
        </StoreProvider>
      </body>
    </html>
  )
}
