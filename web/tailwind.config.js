/* eslint-disable @typescript-eslint/no-var-requires */
const defaultTheme = require('tailwindcss/defaultTheme')

/** @type {import('tailwindcss').Config} */
module.exports = {
  experimental: {
    optimizeUniversalDefaults: true
  },
  content: ['./src/**/*.{js,jsx,ts,tsx,vue,mdx,md}', './content/**/*.{js,jsx,ts,tsx,vue,mdx,md}'],
  darkMode: 'class', // or 'media' or 'class
  theme: {
    screens: {
      xxs: '320px',
      xs: '375px',
      xsm: '425px',
      '3xl': '1920px',
      '4xl': '2560px',
      '5xl': '3840px',
      ...defaultTheme.screens
    },
    container: {
      center: true,
      padding: '2rem',
      screens: {
        '2xl': '1400px'
      }
    },
    extend: {
      fontFamily: {
        sans: ['var(--font-mono)', ...defaultTheme.fontFamily.sans],
        serif: ['var(--font-lobster)', ...defaultTheme.fontFamily.serif],
        mono: ['var(--font-satoshi)', ...defaultTheme.fontFamily.mono]
      },
      fontSize: {
        xs: ['0.65rem', '0.75rem'],
        xsm: ['0.75rem', '1rem']
      },
      maxWidth: {
        '8xl': '90rem',
        '9xl': '100rem',
        '10xl': '110rem',
        '11xl': '120rem'
      },
      transitionProperty: {
        height: 'height',
        width: 'width',
        spacing: 'margin, padding',
        maxHeight: 'max-height',
        maxWidth: 'max-width'
      }
    }
  },
  plugins: []
}
