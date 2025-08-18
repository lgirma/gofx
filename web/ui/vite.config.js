import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'
import { resolve } from 'path';

// https://vite.dev/config/
export default defineConfig(({command}) => {
  return {
    base: command === "build" ? "/app/" : "/",
    plugins: [svelte()],
    resolve: {
      alias: {
        $lib: resolve("./src/lib"),
      },
    },
    build: {
      lib: {
        entry: resolve("./src/index.js"),
        name: "gofx-web-ui",
      },
      rollupOptions: {
        external: ['svelte'],
        output: {
          globals: {
            svelte: 'svelte',
          }
        },
      },
    },
  }
});
