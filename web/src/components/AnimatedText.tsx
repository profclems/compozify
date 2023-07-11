'use client'

import { motion } from 'framer-motion'
import { cn } from '~/utils/classNames'

export default function AnimatedText({
  text,
  className,
  ...props
}: {
  text: string
  className?: string
  props?: React.HTMLAttributes<HTMLElement>
}) {
  const words = text.split('')

  const container = {
    hidden: { opacity: 0 },
    visible: (i: number) => ({
      opacity: 1,
      transition: {
        staggerChildren: 0.1,
        delayChildren: i * 0.05
      }
    })
  }

  const child = {
    visible: {
      opacity: 1,
      y: 0,
      transition: {
        type: 'spring',
        damping: 12,
        stiffness: 100
      }
    },
    hidden: {
      opacity: 0,
      y: 20,
      transition: {
        type: 'spring',
        damping: 20,
        stiffness: 100
      }
    }
  }

  return (
    <motion.h1
      className={cn('flex overflow-hidden', className)}
      variants={container}
      initial="hidden"
      animate="visible"
      {...props}
    >
      {words.map((char, index) => (
        <motion.span
          key={index}
          className="inline-block"
          variants={child}
          initial="hidden"
          animate="visible"
        >
          {char}
        </motion.span>
      ))}
    </motion.h1>
  )
}
