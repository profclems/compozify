'use client'

import Link, { LinkProps } from 'next/link'
import { useRouter } from 'next/navigation'
import DocsSideNav from '~/components/docs-side-nav'
import { buttonVariants } from '~/components/ui/button'
import { ScrollArea } from '~/components/ui/scroll-area'
import { Sheet, SheetContent, SheetTrigger } from '~/components/ui/sheet'
import { docsConfig } from '~/config/docs'
import { siteConfig } from '~/config/site'
import useStore from '~/context/useStore'
import { cn } from '~/utils/classNames'
import { HiMenu } from 'react-icons/hi'

export default function MobileSideNav() {
  const { menu, setMenu } = useStore()

  return (
    <Sheet open={menu} onOpenChange={setMenu}>
      <SheetTrigger asChild>
        <button type="button" className={cn(buttonVariants({ variant: 'ghost' }), 'md:hidden')}>
          <HiMenu className="h-5 w-auto" />
          <span className="sr-only">Toggle Menu</span>
        </button>
      </SheetTrigger>
      <SheetContent side="left" className="bg-white/95 pr-0 dark:bg-zinc-800/95">
        <MobileLink href="/" className="flex items-center" onOpenChange={setMenu}>
          <span className="font-bold uppercase">{siteConfig.name}</span>
        </MobileLink>
        <ScrollArea className="my-4 h-[calc(100vh-8rem)] pb-10 pl-2">
          <DocsSideNav items={docsConfig} />
        </ScrollArea>
      </SheetContent>
    </Sheet>
  )
}

interface MobileLinkProps extends LinkProps {
  onOpenChange?: (open: boolean) => void
  children: React.ReactNode
  className?: string
}

function MobileLink({ href, onOpenChange, className, children, ...props }: MobileLinkProps) {
  const router = useRouter()
  return (
    <Link
      href={href}
      onClick={() => {
        router.push(href.toString())
        onOpenChange?.(false)
      }}
      className={cn(className)}
      {...props}
    >
      {children}
    </Link>
  )
}
