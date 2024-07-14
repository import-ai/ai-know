<template>
  <div v-html="htmlContent"></div>
</template>

<script setup>
import { ref, onMounted } from 'vue';

const htmlContent = ref('');

onMounted(async () => {
  try {
    const response = await fetch('/chat.html');
    if (!response.ok) {
      throw new Error('Network response was not ok');
    }
    htmlContent.value = await response.text();
  } catch (error) {
    console.error('Failed to load HTML content:', error);
    htmlContent.value = '<p>Error loading content.</p>';
  }
});
</script>
