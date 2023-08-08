import CopyButton from '~/components/copy-button'
import { LiaFileInvoiceSolid } from 'react-icons/lia'
import hljs from 'highlight.js'
import { cn } from '~/utils/classNames'
import 'highlight.js/styles/atom-one-dark.css'
import { stripHtml } from '~/utils/stripHtml'

export default function PreviewEditor({ className, code }: { className?: string; code: string }) {
  const highlightedCode = hljs.highlight(code || '', { language: 'yaml' }).value

  return (
    <div className={cn('my-4 rounded border border-neutral-900 dark:border-neutral-800', className)}>
      {/* header */}
      <div className="flex items-center justify-between bg-neutral-800 px-4 py-2 dark:bg-neutral-900">
        <h2 className="flex max-w-[80%] items-center space-x-2">
          <LiaFileInvoiceSolid className="h-4 w-auto text-white" aria-hidden />
          <span className="text-neutral-400">docker-compose.yaml</span>
        </h2>
        <CopyButton
          value={stripHtml(highlightedCode)}
          className={cn('border-none text-white opacity-50 hover:bg-transparent hover:opacity-100')}
        />
      </div>
      <pre
        className="overflow-x-auto bg-neutral-900 !font-sans text-white px-2 py-4 dark:bg-black sm:px-4"
        dangerouslySetInnerHTML={{ __html: highlightedCode }}
      />
    </div>
  )
}
