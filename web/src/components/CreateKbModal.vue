<script setup>
const props = defineProps({
  isOpen: Boolean
});
const emit = defineEmits(['close', 'submit']);

import { ref, watch } from 'vue';

const newKbTitle = ref('');

watch(props, (newProps) => {
  if (!newProps.isOpen) {
    newKbTitle.value = '';
  }
});
</script>

<template>
  <div v-if="props.isOpen" class="fixed inset-0 bg-gray-600 bg-opacity-50 flex items-center justify-center">
    <div class="bg-white p-4 rounded shadow-lg w-1/3">
      <h2 class="text-xl mb-4">Create New Knowledge Base</h2>
      <input v-model="newKbTitle" type="text" placeholder="Enter knowledge base title"
        class="w-full p-2 border rounded mb-4" />
      <div class="flex justify-end">
        <button @click="emit('close')"
          class="bg-gray-500 text-white py-2 px-4 rounded mr-2 hover:bg-gray-600">Cancel</button>
        <button @click="emit('submit', newKbTitle)"
          class="bg-blue-700 text-white py-2 px-4 rounded hover:bg-blue-800">Create</button>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
