'use client'

import { useEffect, useRef } from 'react'
import { useInView } from 'framer-motion'
import { useForm, SubmitHandler } from 'react-hook-form'
import { FaDocker } from 'react-icons/fa'
import { cn } from '~/utils/classNames'
import { siteConfig } from '~/config/site'
import useStore from '~/context/useStore'

export default function Home() {
  const { titleInView: t, setTitleInView } = useStore()
  const titleRef = useRef<HTMLHeadingElement>(null)
  const titleInView = useInView(titleRef, { margin: '0px 0px -100px 0px' })
  const commandInputRef = useRef<HTMLInputElement>(null)

  useEffect(() => {
    if (titleInView !== t) setTitleInView(titleInView)
  }, [setTitleInView, t, titleInView])

  useEffect(() => {
    function onKeyDown(e: KeyboardEvent) {
      if (
        e.key !== '/' ||
        (e.target as HTMLElement).tagName === 'INPUT' ||
        (e.target as HTMLElement).tagName === 'SELECT' ||
        (e.target as HTMLElement).tagName === 'TEXTAREA' ||
        (e.target as HTMLElement).isContentEditable
      ) {
        return
      }
      e.preventDefault()
      commandInputRef.current?.focus()
    }
    window.addEventListener('keydown', onKeyDown)
    return () => window.removeEventListener('keydown', onKeyDown)
  }, [])

  interface CommandInput {
    command: string
  }

  const {
    register,
    handleSubmit,
    formState: { errors }
  } = useForm<CommandInput>({
    mode: 'onChange',
    reValidateMode: 'onChange'
  })

  const onSubmit: SubmitHandler<CommandInput> = _data => {}

  return (
    <main className={cn('')}>
      {/* header */}
      <header className="min-h-[40vh] sm:min-h-[35vh] flex justify-center flex-col px-5 sm:px-8 lg:px-16">
        <div className="">
          <div className="space-y-8 max-w-lg mx-auto">
            <h1 ref={titleRef} className="font-bold text-4xl text-center uppercase">
              {siteConfig.name}
            </h1>
            <p className="text-center text-lg">{siteConfig.description}</p>
          </div>
        </div>
      </header>
      {/* main */}
      <main className="px-5 sm:px-8 lg:px-16">
        <form onSubmit={handleSubmit(onSubmit)}>
          <div className="px-2 max-w-lg mx-auto space-y-4">
            {/* label */}
            <label htmlFor="command-input" className="font-bold">
              Enter Command to Generate `<em>docker-compose.yml</em>` *
            </label>
            {/* input */}
            <div className={cn('flex flex-col gap-5')}>
              <input
                type="text"
                id="command-input"
                autoComplete="off"
                {...register('command', {
                  required: true,
                  validate: value => value.startsWith('docker run')
                })}
                ref={commandInputRef}
                placeholder={`(Press “/” to focus)`}
                className={cn(
                  'flex-auto border bg-white border-zinc-500 focus:border-zinc-950 py-2 px-4 text-zinc-950 placeholder-zinc-500 focus:placeholder-zinc-400 focus:outline-none focus:ring-0 dark:bg-zinc-600/60 dark:text-zinc-200 dark:border-zinc-400 dark:focus:border-zinc-50'
                )}
              />
              <button
                type="submit"
                className="bg-zinc-950 uppercase flex items-center justify-center transition-transform hover:-translate-y-0.5 text-white space-x-4 px-4 py-3"
              >
                <FaDocker className="h-4 w-auto" />
                <span className="">Generate</span>
              </button>
            </div>
            {/* errors */}
            {errors.command && (
              <div className="">
                <p className="text-red-500">{errors.command.message}</p>
              </div>
            )}
          </div>
        </form>
      </main>
    </main>
  )
}
