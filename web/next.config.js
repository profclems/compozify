/* eslint-disable @typescript-eslint/no-var-requires */
const path = require('path')
const { withContentlayer } = require('next-contentlayer')
const runtimeCaching = require('next-pwa/cache')
const withPWA = require('next-pwa')({
  dest: 'public',
  runtimeCaching,
  disable: process.env.NODE_ENV === 'development'
})

/** @type {import('next').NextConfig} */
const nextConfig = {
  webpack: (config, { defaultLoaders }) => {
    // clear cache
    defaultLoaders.babel.options.cache = false

    // resolve path
    config.resolve.modules.push(path.resolve(`./`))

    return config
  },
  async redirects() {
    return require('./redirects.json')
  }
}

const withALL = (nextConfig = {}) => withContentlayer(withPWA(nextConfig))

module.exports = withALL(nextConfig)
