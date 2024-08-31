/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    'internal/templates/*.templ',
    'static/**/*.{html,js,css}',
  ],
  theme: {
    extend: {
      fontFamily :{
        merienda: ["Merienda", "cursive"], 
      },
      backgroundImage :{
        'Almeria': "url('/static/images/background.png')"
      }
    },
  },
  plugins: [],
}

