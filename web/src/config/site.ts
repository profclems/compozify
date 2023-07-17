interface SiteConfig {
  name: string
  url: string
  description: {
    short: string
    long: string
  }
  links: {
    github: string
  }
}

export const siteConfig: SiteConfig = {
  name: 'compozify',
  url: 'https://github.com/profclems/compozify',
  description: {
    short:
      'Compozify is a simple (yet complicated) tool to generate a docker-compose.yml file from a docker run command.',
    long: 'Say goodbye to complex `docker run` commands and embrace the power of `$ docker-compose up!` Compozify generates a clean and concise `docker-compose.yml` file for you from your docker run commands.'
  },

  links: {
    github: 'https://github.com/profclems/compozify'
  }
}
