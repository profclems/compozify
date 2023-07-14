import '~/styles/mdx.css'
import { notFound } from 'next/navigation'
import { allDocs } from 'contentlayer/generated'
import { getTableOfContents } from '~/lib/toc'
import { Mdx } from '~/components/mdx'
import { cn } from '~/utils/classNames'
import { DocsPaginate as Paginate } from '~/components/paginate'
import { DocsTableOfContents } from '~/components/toc'
import { HiChevronRight } from 'react-icons/hi'
import { ScrollArea } from '~/components/ui/scroll-area'

interface DocPageProps {
  params: {
    slug: string[]
  }
}

async function getDocFromParams({ params }: DocPageProps) {
  const slug = params.slug?.join('/') || ''
  const doc = allDocs.find(doc => doc.slugAsParams === slug)

  if (!doc) {
    null
  }

  return doc
}

export async function generateStaticParams(): Promise<DocPageProps['params'][]> {
  return allDocs.map(doc => ({ slug: doc.slugAsParams.split('/') }))
}

export default async function DocPage({ params }: DocPageProps) {
  const doc = await getDocFromParams({ params })

  if (!doc) {
    return notFound()
  }

  const toc = await getTableOfContents(doc?.body.raw)

  return (
    <main className="relative py-6 lg:gap-10 lg:py-8 xl:grid xl:grid-cols-[1fr_300px]">
      <div className="mx-auto w-full min-w-0">
        <div className="mb-4 flex items-center space-x-1 text-sm">
          <div className="overflow-hidden text-ellipsis whitespace-nowrap">Docs</div>
          <HiChevronRight className="h-4 w-4" />
          <div className="font-medium text-foreground">{doc.title}</div>
        </div>
        <div className="pb-12 pt-4">
          <Mdx code={doc.body.code} />
        </div>
        <Paginate docs={allDocs} activeDocs={doc} />
      </div>
      {toc && (
        <div className="hidden text-sm xl:block">
          <div className="sticky top-16 -mt-8 h-[calc(100vh-3.5rem)] overflow-hidden pt-12">
            <ScrollArea className="pb-10">
              <DocsTableOfContents toc={toc} />
            </ScrollArea>
          </div>
        </div>
      )}
    </main>
  )
}
