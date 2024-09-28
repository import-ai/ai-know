import { defineConfig } from 'orval';

// http://ddns.ip.lucien.ink:8911/swagger/doc.json
export default defineConfig({
  'ai-know': {
    output: {
      mode: 'tags-split',
      target: './src/api.ts',
      schemas: 'src/model',
      client: 'swr',
      mock: true,
    },
    input: {
      target: './doc.json',
    },
  },
});
