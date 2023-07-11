/* eslint-disable @typescript-eslint/no-var-requires */
const path = require('path')
const runtimeCaching = require('next-pwa/cache')
const withPWA = require('next-pwa')({
  dest: 'public',
  runtimeCaching,
  disable: process.env.NODE_ENV === 'development'
})

/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  images: { disableStaticImages: true },
  webpack: (config, { defaultLoaders }) => {
    // clear cache
    defaultLoaders.babel.options.cache = false

    // resolve path
    config.resolve.modules.push(path.resolve(`./`))

    return config
  }
}

module.exports = withPWA(nextConfig)
