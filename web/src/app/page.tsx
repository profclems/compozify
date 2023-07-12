'use client'

import { useEffect, useRef } from 'react'
import { useInView } from 'framer-motion'
import { cn } from '~/utils/classNames'
import { siteConfig } from '~/config/site'
import useStore from '~/context/useStore'

export default function Home() {
  const { titleInView: t, setTitleInView } = useStore()
  const titleRef = useRef<HTMLHeadingElement>(null)
  const titleInView = useInView(titleRef, { margin: '0px 0px -100px 0px' })

  useEffect(() => {
    if (titleInView !== t) setTitleInView(titleInView)
  }, [setTitleInView, t, titleInView])

  return (
    <main className={cn('')}>
      {/* header */}
      <header className="min-h-[40vh] sm:min-h-[35vh] flex justify-center flex-col">
        <div className="px-5 sm:px-8 lg:px-16">
          <div className="space-y-8 max-w-lg mx-auto">
            <h1 ref={titleRef} className="font-bold text-4xl text-center uppercase">
              {siteConfig.name}
            </h1>
            <p className="text-center text-lg">{siteConfig.description}</p>
          </div>
        </div>
      </header>
      {/* main */}
      <main className=""></main>
    </main>
  )
}
