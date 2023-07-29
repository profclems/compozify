import { getHighlighter, Lang } from 'shiki'
import { languages } from '~/lib/rehype-languages'

const highlighterPromise = getHighlighter({
  theme: 'nord',
  langs: languages as Lang[]
})

export async function hightlight(code: string, language: string) {
  const highlighter = await highlighterPromise
  const output = highlighter.codeToHtml(code, { lang: language })
  return output
}
