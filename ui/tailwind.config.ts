import type { Config } from 'tailwindcss'

const {nextui} = require("@nextui-org/react");
const defaultTheme = require('tailwindcss/defaultTheme')

const config: Config = {
  content: [
    './src/pages/**/*.{js,ts,jsx,tsx,mdx}',
    './src/components/**/*.{js,ts,jsx,tsx,mdx}',
    './src/app/**/*.{js,ts,jsx,tsx,mdx}',
    "./node_modules/@nextui-org/theme/dist/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      maxWidth: {
        '1/4': '25%',
        '3/10': '30%',
        '2/5': '40%',
        '1/2': '50%',
        '3/5': '60%',
        '7/10': '70%',
        '3/4': '75%',
        '4/5': '80%',
        '9/10': '90%',

      }
      // backgroundImage: {
      //   'gradient-radial': 'radial-gradient(var(--tw-gradient-stops))',
      //   'gradient-conic':
      //     'conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))',
      // },
    },
  },
  fontFamily: {
    'sans': ['"Proxima Nova"', ...defaultTheme.fontFamily.sans],
  },
  darkMode: "class",
  plugins: [nextui()],
}
export default config
