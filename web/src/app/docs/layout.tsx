import { ReactNode } from 'react'
import { docsConfig } from '~/config/docs'
import { ScrollArea } from "~/components/ui/scroll-area"
import DocsSideNav from '~/components/docs-side-nav'

export default function Layout({ children }: { children: ReactNode }) {
  return (
    <div className="border-b">
      <div className="container flex-1 items-start md:grid md:grid-cols-[240px_minmax(0,1fr)] md:gap-6 lg:grid-cols-[275px_minmax(0,1fr)] lg:gap-10">
        <aside className="fixed top-16 z-30 -ml-2 hidden h-[calc(100vh-3.5rem)] w-full shrink-0 md:sticky md:block">
          <ScrollArea className="h-full py-6 pl-8 pr-6 lg:py-8">
            <DocsSideNav items={docsConfig} />
          </ScrollArea>
        </aside>
        {children}
      </div>
    </div>
  )
}
