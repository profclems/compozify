export default function Custom404() {
  return (
    <section className={'fixed inset-0 grid h-screen place-content-center px-4 text-center'}>
      <div className="">
        <div className="flex flex-auto flex-col items-center justify-center sm:flex-row">
          <h1 className="text-7xl font-extrabold tracking-tight text-zinc-900 dark:text-white/40 sm:mr-6 sm:border-r sm:border-zinc-900/50 sm:pr-6 sm:dark:border-white/50">
            404
          </h1>
          <h2 className="mt-2 text-lg text-zinc-700 dark:text-white sm:mt-0">This page could not be found.</h2>
        </div>
      </div>
    </section>
  )
}
