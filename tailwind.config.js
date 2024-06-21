/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./view/**/*.templ"],
  theme: {
    extend: {
      backgroundImage: {
        'admin-image': "url('/public/images/abstract-paper-background/509434-PITAEO-611.eps')"
      }
    },
  },
  plugins: [],
}

