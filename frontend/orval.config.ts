import { defineConfig } from 'orval';

// http://ddns.ip.lucien.ink:8911/swagger/doc.json
export default defineConfig({
  'ai-know': {
    output: {
      mode: 'tags-split',
      workspace: 'src/api',
      target: '',
      schemas: 'model',
      client: 'swr',
      httpClient:'axios',
      mock: false,
    },
    input: {
      target: './doc.json',
    },
  },
});
