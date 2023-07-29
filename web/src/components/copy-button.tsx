'use client'

import { useEffect, useState } from 'react'
import { cn } from '~/utils/classNames'
import copyToClipboardWithMeta from '~/utils/copy'
import { HiCheck, HiClipboardCopy } from 'react-icons/hi'

interface CopyButtonProps extends React.HTMLAttributes<HTMLButtonElement> {
  value: string
  src?: string
}

export default function CopyButton({ value, className, src, ...props }: CopyButtonProps) {
  const [hasCopied, setHasCopied] = useState(false)

  useEffect(() => {
    setTimeout(() => {
      setHasCopied(false)
    }, 2000)
  }, [hasCopied])

  return (
    <button
      className={cn(
        'relative z-20 inline-flex h-8 items-center justify-center rounded-md border-neutral-200 p-2 text-sm font-medium text-neutral-900 transition-all hover:bg-neutral-100 focus:outline-none dark:text-neutral-100 dark:hover:bg-neutral-800',
        className
      )}
      onClick={() => {
        copyToClipboardWithMeta(value, {
          component: src
        })
        setHasCopied(true)
      }}
      {...props}
    >
      <span className="sr-only">Copy</span>
      {hasCopied ? <HiCheck className="h-4 w-auto" /> : <HiClipboardCopy className="h-4 w-auto" />}
    </button>
  )
}
