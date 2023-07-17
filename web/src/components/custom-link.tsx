'use client'

import { forwardRef, ReactNode } from 'react'
import Link, { LinkProps } from 'next/link'
import { cn } from '~/utils/classNames'

export interface CustomLinkProps extends Omit<React.AnchorHTMLAttributes<HTMLAnchorElement>, 'href'>, LinkProps {
  className?: string
  children: ReactNode
}

const CustomLink = forwardRef<HTMLAnchorElement, CustomLinkProps>(({ className, children, ...props }, ref) => {
  return (
    <Link ref={ref} className={cn('group relative cursor-pointer px-1 py-0.5', className)} {...props}>
      {children}
      <span className="absolute inset-x-0 bottom-0 block bg-zinc-900 h-[3px] w-0 rounded-md transition-width group-hover:w-full dark:bg-white" />
    </Link>
  )
})

export default CustomLink
