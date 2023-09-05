/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      backgroundImage: {
        "gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
        "gradient-conic":
          "conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))",
      },
      width: {
        128: "32rem",
      },
      screens: {
        midsm: "30rem",
        mid: "48rem",
        midlg: "82rem",
      },
      colors: {
        primary: {
          DEFAULT: "rgb(var(--color-primary-rgb))",
          transparent: "rgba(var(--color-primary-rgb), 0.15)",
        },
        secondary: "rgb(var(--color-secondary-rgb))",
        gray: {
          DEFAULT: "rgb(var(--color-gray-rgb))",
          light: "rgb(var(--color-gray-light-rgb))",
          lightest: "rgb(var(--color-gray-lightest-rgb))",
          dark: "rgb(var(--color-gray-dark-rgb))",
        },
      },
      fontSize: {
        tiny: ".625rem",
      },
      animation: {
        "slide-bottom": "slide-bottom 0.5s both",
      },
      keyframes: {
        "slide-bottom": {
          "0%": {
            transform: "translateY(0)",
          },
          to: {
            transform: "translateY(5px)",
          },
        },
      },
    },
  },
  plugins: [],
}
