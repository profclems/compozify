'use client'

import { useEffect, useMemo, useState } from 'react'
import Link from 'next/link'
import { ScrollToTopWithDocs } from '~/components/scroll-to-top'
import { cn } from '~/utils/classNames'
import clsx from 'clsx'
import { TableOfContents } from 'lib/toc'

interface TocProps {
  toc: TableOfContents
  className?: string
}

export function DocsTableOfContents({ toc, className }: TocProps) {
  const itemIds = useMemo(() => {
    if (toc && toc.items) {
      return toc.items
        .flatMap(item => [item.url, item?.items?.map(i => i.url)] || [])
        .flat()
        .filter(Boolean)
        .map(id => id?.split('#')[1])
    }

    return []
  }, [toc])
  const activeHeading = useActiveItem(itemIds as string[])

  if (!toc?.items) {
    return null
  }

  return (
    <aside className={cn('text-xs transition-transform xl:text-sm', className)}>
      {/* Table of content */}
      <div className={cn('sticky top-16 -mt-10 max-h-[calc(var(--vh)-4rem)] space-y-2 overflow-y-auto pr-2 pt-16')}>
        <p className="text-sm font-medium uppercase">On This Page</p>
        {/* scroll to top */}
        <ScrollToTopWithDocs />
        <Tree tree={toc} activeItem={activeHeading} />
      </div>
    </aside>
  )
}

function useActiveItem(itemIds: string[]) {
  const [activeId, setActiveId] = useState<string | undefined>(undefined)

  useEffect(() => {
    const observer = new IntersectionObserver(
      entries => {
        entries.forEach(entry => {
          if (entry.isIntersecting) {
            setActiveId(entry.target.id || '')
          }
        })
      },
      { rootMargin: `0% 0% -80% 0%` }
    )

    itemIds?.forEach(id => {
      const element = document.getElementById(id)
      if (element) observer.observe(element)
    })

    return () => {
      itemIds?.forEach(id => {
        const element = document.getElementById(id)
        if (element) observer.unobserve(element)
      })
    }
  }, [itemIds])

  return activeId
}

interface TreeProps {
  tree: TableOfContents
  level?: number
  activeItem?: string
}

function Tree({ tree, level = 1, activeItem }: TreeProps) {
  return tree?.items?.length && level < 4 ? (
    <ul
      className={cn('m-0 list-none', { 'pl-1.5': level !== 1 })}
      style={{
        paddingLeft: `${level * 0.5}rem`
      }}
    >
      {tree.items.map((item, index) => {
        return (
          <li key={index} className={clsx('mt-0 pt-2')}>
            <a
              href={item.url}
              className={cn(
                'inline-block text-sm no-underline',
                item.url === `#${activeItem}`
                  ? 'font-medium text-rose-600 dark:text-orange-300'
                  : 'text-neutral-700 hover:text-neutral-900 dark:text-neutral-400'
              )}
            >
              {item.title}
            </a>
            {item.items?.length ? <Tree tree={item} level={level + 1} activeItem={activeItem} /> : null}
          </li>
        )
      })}
    </ul>
  ) : null
}
