import { DocsConfig } from '~/types/nav'

export const docsConfig: DocsConfig = {
  'Getting Started': [
    {
      title: 'Try it out',
      href: '/'
    },
    {
      title: 'GitHub',
      href: 'https://github.com/profclems/compozify',
      external: true
    },
    {
      title: 'Installation',
      href: '/docs/installation'
    }
  ],
  Commands: [
    {
      title: 'compozify',
      href: '/docs/commands/compozify'
    },
    {
      title: 'add-service',
      href: '/docs/commands/compozify-add-service'
    },
    {
      title: 'convert',
      href: '/docs/commands/compozify-convert'
    }
  ]
}
