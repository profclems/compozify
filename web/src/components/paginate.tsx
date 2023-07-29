'use client'

import Link from 'next/link'
import { buttonVariants } from '~/components/ui/button'
import useMedia from '~/hooks/useMedia'
import { cn } from '~/utils/classNames'
import { Docs } from 'contentlayer/generated'
import { HiChevronLeft, HiChevronRight } from 'react-icons/hi'

interface DocsPaginationProps {
  docs: Docs[]
  activeDocs: Docs
}

export function DocsPaginate({ docs, activeDocs }: DocsPaginationProps) {
  const pager = getPagerForDocs(docs, activeDocs)
  const xsm = useMedia('(max-width: 425px)')

  return (
    <div className="relative flex flex-row items-center justify-between py-5">
      {pager && pager.prev?.slug && (
        <Link href={pager.prev.slug} className={cn(buttonVariants({ variant: 'outline' }), 'absolute left-0')}>
          <HiChevronLeft className="mr-2 h-4 w-4" />
          <span className="overflow-hidden text-ellipsis whitespace-nowrap max-lg:max-w-[10rem] max-sm:text-xsm">
            {xsm ? 'Previous' : pager.prev.title}
          </span>
        </Link>
      )}
      {pager && pager.next?.slug && (
        <Link href={pager.next.slug} className={cn(buttonVariants({ variant: 'outline' }), 'absolute right-0')}>
          <span className="overflow-hidden text-ellipsis whitespace-nowrap max-lg:max-w-[10rem] max-md:text-xsm">
            {xsm ? 'Next' : pager.next.title}
          </span>
          <HiChevronRight className="ml-2 h-4 w-4" />
        </Link>
      )}
    </div>
  )
}

export function getPagerForDocs(docs: Docs[], activeDocs: Docs) {
  const flattenedLinks = [null, ...docs, null]
  const activeIndex = flattenedLinks.findIndex(link => activeDocs.slug === link?.slug)
  const prev = activeIndex !== 0 ? flattenedLinks[activeIndex - 1] : null
  const next = activeIndex !== flattenedLinks.length - 1 ? flattenedLinks[activeIndex + 1] : null
  return {
    prev,
    next
  }
}
