export interface NavItem {
  title: string
  href?: string
  disabled?: boolean
  external?: boolean
  label?: string
  icon?: React.ReactNode
  children?: NavItem[]
}

export interface DocsConfig {
  'Getting Started': NavItem[]
  Commands: NavItem[]
}

export interface ErrorCause extends Error {
  cause?: { error: Error; res: Response }
}
