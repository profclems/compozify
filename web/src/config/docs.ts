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
      title: 'Compozify',
      href: '/docs/commands/compozify'
    },
    {
      title: 'Convert',
      href: '/docs/commands/compozify-convert'
    },
    {
      title: 'Add Service',
      href: '/docs/commands/compozify-add-service'
    }
  ]
}
