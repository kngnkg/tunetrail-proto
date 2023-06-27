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
    },
  },
  plugins: [],
}
