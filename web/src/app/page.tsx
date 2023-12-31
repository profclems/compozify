'use client'

import { Fragment, useEffect, useRef, useState } from 'react'
import Spinner from '~/components/Spinner'
import { siteConfig } from '~/config/site'
import useStore from '~/context/useStore'
import { cn } from '~/utils/classNames'
import { useInView } from 'framer-motion'
import { SubmitHandler, useForm } from 'react-hook-form'
import { FaDocker } from 'react-icons/fa'
import CopyButton from '~/components/copy-button'
import { LiaFileInvoiceSolid } from 'react-icons/lia'
import CustomLink from '~/components/custom-link'

export default function Home() {
  const { titleInView: t, setTitleInView, compose, code, error } = useStore()
  const titleRef = useRef<HTMLHeadingElement>(null)
  const titleInView = useInView(titleRef, { margin: '0px 0px -100px 0px' })
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    if (titleInView !== t) setTitleInView(titleInView)
  }, [setTitleInView, t, titleInView])

  interface CommandInput {
    command: string
  }

  const {
    register,
    handleSubmit,
    setFocus,
    formState: { errors, isSubmitting, isValid }
  } = useForm<CommandInput>({ mode: 'onChange', reValidateMode: 'onChange' })

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
      setFocus('command')
    }
    window.addEventListener('keydown', onKeyDown)
    return () => window.removeEventListener('keydown', onKeyDown)
  }, [setFocus])

  useEffect(() => {
    if (isSubmitting && isValid) setLoading(true)
    else setLoading(false)
  }, [isSubmitting, isValid])

  const onSubmit: SubmitHandler<CommandInput> = data => compose(data.command).finally(() => setLoading(false))

  return (
    <main className={cn('space-y-8')}>
      {/* header */}
      <header className="flex flex-col justify-center px-5 sm:px-8 lg:px-16">
        <div className="">
          <div className="mx-auto max-w-5xl space-y-8">
            <h1 ref={titleRef} className="text-center text-4xl font-bold">
              $ {siteConfig.name}
            </h1>
            <p className="text-center text-lg">
              Say goodbye to complex{' '}
              <CustomLink
                href="https://docs.docker.com/engine/reference/run/"
                target="_blank"
                rel="noreferrer noopener"
              >
                <em className="px-0.5 bg-zinc-500/20 font-bold">docker run</em>
              </CustomLink>
              commands and embrace the power of{' '}
              <em className="px-0.5 bg-zinc-300/20 font-bold">$ docker-compose up!</em>. Compozify generates a clean and
              concise <em className="px-0.5 bg-zinc-300/20 font-bold">docker-compose.yml</em> file for you from your
              docker run commands.
            </p>
          </div>
        </div>
      </header>
      {/* main */}
      <div className="px-5 sm:px-8 lg:px-16">
        {/* form */}
        <form action="#" method="post" onSubmit={handleSubmit(onSubmit)} className="">
          <div className="mx-auto max-w-5xl space-y-4 px-2">
            {/* label */}
            <div className={cn('flex flex-col gap-5')}>
              <label htmlFor="command-input" className="">
                Install{' '}
                <CustomLink
                  href="https://github.com/profclems/compozify#installation"
                  target="_blank"
                  rel="noreferrer noopener"
                  className="px-1 bg-zinc-300/20 font-bold"
                >
                  the CLI
                </CustomLink>{' '}
                or Paste your{' '}
                <CustomLink
                  href="https://docs.docker.com/engine/reference/run/"
                  target="_blank"
                  rel="noreferrer noopener"
                  className="px-1 bg-zinc-300/20 font-bold"
                >
                  docker run
                </CustomLink>{' '}
                command below!
              </label>
              {/* input */}
              <input
                type="text"
                id="command-input"
                autoComplete="off"
                {...register('command', {
                  required: 'This field is required',
                  pattern: {
                    value: /^docker run .+$/i,
                    message: 'Enter a valid docker command starting with “docker run”'
                  }
                })}
                placeholder={`(Press “/” to focus)`}
                className={cn(
                  'flex-auto border border-zinc-500 bg-white px-4 py-2 text-zinc-950 placeholder-zinc-500 focus:border-zinc-950 focus:placeholder-zinc-400 focus:outline-none focus:ring-0 dark:border-zinc-400 dark:bg-zinc-600/60 dark:text-zinc-200 dark:focus:border-zinc-50'
                )}
              />
              <button
                type="submit"
                className="flex items-center justify-center space-x-4 bg-zinc-950 px-4 py-3 text-white transition-transform hover:-translate-y-0.5"
                disabled={loading}
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
            {errors.command && <p className="text-red-500">{errors.command.message}</p>}
          </div>
        </form>
      </div>
      {/* data from fetch */}
      <div className="mt-4 px-5 sm:px-8 lg:px-16">
        <div className="mx-auto max-w-5xl">
          {code && (
            <div className="my-4 rounded border border-neutral-900 dark:border-neutral-800">
              {/* header */}
              <div className="flex items-center justify-between bg-neutral-900 px-4 py-2 dark:bg-neutral-800">
                <h2 className="flex max-w-[80%] items-center space-x-2">
                  <LiaFileInvoiceSolid className="h-4 w-auto text-white" aria-hidden />
                  <span className="text-neutral-400">docker-compose.yaml</span>
                </h2>
                <CopyButton
                  value={code.replace(/^\s*\|/, '')}
                  className={cn('border-none text-white opacity-50 hover:bg-transparent hover:opacity-100')}
                />
              </div>
              <pre className="overflow-x-auto bg-neutral-900 !font-sans text-white px-2 py-4 dark:bg-black sm:px-4">
                {code.replace(/^\s*\|/, '')}
              </pre>
            </div>
          )}
          {error && (
            <div className="flex flex-col gap-4">
              <p className="text-red-500">{error.message}</p>
              <p className="text-red-500">Error code: {error.code}</p>
            </div>
          )}
        </div>
      </div>
    </main>
  )
}
