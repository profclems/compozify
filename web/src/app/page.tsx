'use client'

import { Fragment, useEffect, useRef, useState } from 'react'
import Spinner from '~/components/Spinner'
import { siteConfig } from '~/config/site'
import useStore from '~/context/useStore'
import { cn } from '~/utils/classNames'
import { useInView } from 'framer-motion'
import { SubmitHandler, useForm } from 'react-hook-form'
import { FaDocker } from 'react-icons/fa'

export default function Home() {
  const { titleInView: t, setTitleInView } = useStore()
  const titleRef = useRef<HTMLHeadingElement>(null)
  const titleInView = useInView(titleRef, { margin: '0px 0px -100px 0px' })
  const commandInputRef = useRef<HTMLInputElement>(null)
  const [loading, setLoading] = useState(false)

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
    formState: { errors, isValid, isSubmitting }
  } = useForm<CommandInput>({
    mode: 'onChange',
    reValidateMode: 'onChange'
  })

  useEffect(() => {
    if (isSubmitting && isValid) setLoading(true)
    else setLoading(false)
  }, [isSubmitting, isValid])

  const onSubmit: SubmitHandler<CommandInput> = _data => {}

  return (
    <main className={cn('')}>
      {/* header */}
      <header className="flex min-h-[40vh] flex-col justify-center px-5 sm:min-h-[35vh] sm:px-8 lg:px-16">
        <div className="">
          <div className="mx-auto max-w-lg space-y-8">
            <h1 ref={titleRef} className="text-center text-4xl font-bold uppercase">
              {siteConfig.name}
            </h1>
            <p className="text-center text-lg">{siteConfig.description}</p>
          </div>
        </div>
      </header>
      {/* main */}
      <main className="px-5 sm:px-8 lg:px-16">
        <form onSubmit={handleSubmit(onSubmit)}>
          <div className="mx-auto max-w-lg space-y-4 px-2">
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
                  validate: value =>
                    value.startsWith('docker run') || 'Must start with “docker run”'
                })}
                ref={commandInputRef}
                placeholder={`(Press “/” to focus)`}
                className={cn(
                  'flex-auto border border-zinc-500 bg-white px-4 py-2 text-zinc-950 placeholder-zinc-500 focus:border-zinc-950 focus:placeholder-zinc-400 focus:outline-none focus:ring-0 dark:border-zinc-400 dark:bg-zinc-600/60 dark:text-zinc-200 dark:focus:border-zinc-50'
                )}
              />
              <button
                type="submit"
                className="flex items-center justify-center space-x-4 bg-zinc-950 px-4 py-3 uppercase text-white transition-transform hover:-translate-y-0.5"
              >
                {loading ? (
                  <Spinner className="h-4 w-auto" />
                ) : (
                  <Fragment>
                    <FaDocker className="h-4 w-auto" />
                    <span className="">Generate</span>
                  </Fragment>
                )}
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
