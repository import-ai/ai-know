<script setup>
const props = defineProps({
  kbs: Array,
  curKb: Object,
  userName: String
})
const emit = defineEmits(['kbSelected', 'createKbClicked', 'logoutClicked'])

import { ref, onMounted, onUnmounted } from 'vue';

const isDropdownOpen = ref(false);

function toggleDropdown() {
  isDropdownOpen.value = !isDropdownOpen.value
}

function handleClickOutside(event) {
  const dropdown = document.querySelector('.dropdown');
  if (dropdown && !dropdown.contains(event.target)) {
    isDropdownOpen.value = false;
  }
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside);
});

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside);
});

function handleDropdownButtonClick() {
  toggleDropdown()
}

function handleDropdownItemSelect(kb) {
  emit('kbSelected', kb)
  toggleDropdown()
}

</script>

<template>
  <nav class="bg-blue-600 p-2 text-white flex items-center justify-between relative">
    <!-- Dropdown Menu -->
    <div class="flex">
      <div class="dropdown relative">
        <button @click="handleDropdownButtonClick" v-if="props.curKb"
          class="bg-blue-700 text-white py-2 px-4 rounded mr-2 hover:bg-blue-800 focus:outline-none items-center flex">
          Knowledge Base: {{ props.curKb.name }}
          <span class="ml-2 arrow"></span>
        </button>
        <div v-if="isDropdownOpen" class="dropdown-content absolute mt-1 w-48 rounded shadow-lg bg-white text-gray-800">
          <a href="#" v-for="kb in props.kbs" :key="kb.id" @click.prevent="handleDropdownItemSelect(kb)"
            class="block px-4 py-2 hover:bg-gray-100">
            {{ kb.name }}
          </a>
        </div>
      </div>
      <button @click="emit('createKbClicked')"
        class="bg-blue-700 text-white py-2 px-4 rounded hover:bg-blue-800 focus:outline-none">
        Create New KB
      </button>
    </div>
    <!-- User Info and Logout Button -->
    <div class="flex items-center space-x-4">
      <span class="text-white">Logged in as: {{ props.userName }}</span>
      <button @click="emit('logoutClicked')"
        class="bg-red-600 text-white py-1 px-4 rounded hover:bg-red-700">Logout</button>
    </div>
  </nav>
</template>

<style scoped>
.dropdown-content {
  display: block;
  position: absolute;
  top: 100%;
  left: 0;
  background-color: white;
  border: 1px solid #ddd;
  border-radius: 0.375rem;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  z-index: 10;
}

.arrow {
  width: 0;
  height: 0;
  border-left: 7px solid transparent;
  border-right: 7px solid transparent;
  border-top: 7px solid white;
}
</style>
