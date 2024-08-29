import { fileURLToPath, URL } from 'node:url'
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
  server: {
    proxy: {
      // Inoltra tutte le richieste che iniziano con "/api" al server Go
      '/session': {
        target: 'http://localhost:3000', // URL del tuo server API Go
        changeOrigin: true,
        pathRewrite: { '^/session': '' }, // Rimuove il prefisso /api per evitare conflitti
      },
    },
  },
  define: {
    "__API_URL__": JSON.stringify("http://localhost:3000"),
  }
})
