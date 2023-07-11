interface SiteConfig {
  name: string
  url: string
  description: string
  links: {
    github: string
  }
}

export const siteConfig: SiteConfig = {
  name: 'compozify',
  url: 'https://github.com/profclems/compozify',
  description:
    'Compozify is a simple (yet complicated) tool to generate a docker-compose.yml file from a docker run command.',
  links: {
    github: 'https://github.com/profclems/compozify'
  }
}
