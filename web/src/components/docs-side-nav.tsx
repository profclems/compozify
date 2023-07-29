'use client'

import Link from 'next/link'
import { usePathname } from 'next/navigation'
import { DocsConfig, NavItem } from '~/types/nav'
import { cn } from '~/utils/classNames'
import useStore from '~/context/useStore'

export interface DocsSidebarNavProps {
  items: DocsConfig
  className?: string
}

export default function DocsSideNav({ items, className }: DocsSidebarNavProps) {
  const pathname = usePathname()

  return (
    <div className={cn('w-full', className)}>
      {Object.entries(items).map(([key, value], index) =>
        value.length > 0 ? (
          <div key={index} className={cn('pb-4')}>
            <h3 className="text-foreground mb-1 rounded-md px-2 py-1 text-lg font-bold uppercase">{key}</h3>
            {value.length > 0 && <DocsSidebarNavItems items={value} pathname={pathname} />}
          </div>
        ) : undefined
      )}
    </div>
  )
}

export interface DocsSidebarNavItemsProps {
  items: NavItem[]
  pathname: string
}

export function DocsSidebarNavItems({ items, pathname }: DocsSidebarNavItemsProps) {
  const { setMenu } = useStore()

  return items?.length ? (
    <div className="grid grid-flow-row auto-rows-max text-sm">
      {items.map((item, index) =>
        item.href && !item.disabled ? (
          <Link
            key={index}
            href={item.href}
            className={cn(
              'group flex w-full items-center rounded-md border border-transparent px-2 py-1 hover:underline',
              item.disabled && 'cursor-not-allowed opacity-60',
              pathname === item.href ? 'text-foreground font-medium underline' : 'text-muted-foreground'
            )}
            target={item.external ? '_blank' : ''}
            rel={item.external ? 'noreferrer noopener' : ''}
            onClick={() => setMenu(false)}
          >
            {item.title}
            {item.label && (
              <span className="ml-2 rounded-md bg-[#9c7d14] px-1.5 py-0.5 text-xs leading-none text-[#000000] no-underline group-hover:no-underline">
                {item.label}
              </span>
            )}
          </Link>
        ) : (
          <span
            key={index}
            className={cn(
              'text-muted-foreground flex w-full cursor-not-allowed items-center rounded-md p-2 hover:underline',
              item.disabled && 'cursor-not-allowed opacity-60'
            )}
          >
            {item.title}
            {item.label && (
              <span className="bg-muted text-muted-foreground ml-2 rounded-md px-1.5 py-0.5 text-xs leading-none no-underline group-hover:no-underline">
                {item.label}
              </span>
            )}
          </span>
        )
      )}
    </div>
  ) : null
}
