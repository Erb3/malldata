const animate = require('tailwindcss-animate');

/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: ['class'],

  content: [
    './index.html',
    './src/**/*.{vue,js,ts,jsx,tsx}',
    './frontend/index.html',
    './frontend/src/**/*.{vue,js,ts,jsx,tsx}',
  ],

  theme: {
    container: {
      center: true,
      padding: '2rem',
      screens: {
        '2xl': '1400px',
      },
    },
    extend: {
      colors: {
        border: 'var(--color-border)',
        input: 'var(--color-border)',
        ring: 'var(--color-border)',
        background: 'var(--color-background)',
        foreground: 'var(--color-raised)',
        primary: {
          DEFAULT: 'var(--primary)',
          foreground: 'var(--color-text)',
        },
        secondary: {
          DEFAULT: 'var(--secondary)',
          foreground: 'var(--color-text)',
        },
        destructive: {
          DEFAULT: 'var(--red)',
          foreground: 'var(--color-text)',
        },
        muted: {
          DEFAULT: 'var(--color-mute)',
          foreground: 'var(--color-text)',
        },
        accent: {
          DEFAULT: 'var(--primary)',
          foreground: 'var(--color-text)',
        },
        popover: {
          DEFAULT: 'var(--color-raised)',
          foreground: 'var(--color-text)',
        },
        card: {
          DEFAULT: 'var(--color-raised)',
          foreground: 'var(--color-text)',
        },
      },
      borderRadius: {
        lg: 'var(--radius)',
        md: 'calc(var(--radius) - 2px)',
        sm: 'calc(var(--radius) - 4px)',
      },
      keyframes: {
        'accordion-down': {
          from: { height: 0 },
          to: { height: 'var(--radix-accordion-content-height)' },
        },
        'accordion-up': {
          from: { height: 'var(--radix-accordion-content-height)' },
          to: { height: 0 },
        },
        'collapsible-down': {
          from: { height: 0 },
          to: { height: 'var(--radix-collapsible-content-height)' },
        },
        'collapsible-up': {
          from: { height: 'var(--radix-collapsible-content-height)' },
          to: { height: 0 },
        },
      },
      animation: {
        'accordion-down': 'accordion-down 0.2s ease-out',
        'accordion-up': 'accordion-up 0.2s ease-out',
        'collapsible-down': 'collapsible-down 0.2s ease-in-out',
        'collapsible-up': 'collapsible-up 0.2s ease-in-out',
      },
    },
  },
  plugins: [animate],
};
