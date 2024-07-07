<script setup>
import { ref } from 'vue'
import { useToast } from "vue-toastification";
import { useRouter } from 'vue-router';

import { HTTPError } from '@/http_client'
import { api_client } from '@/api_client'

const toast = useToast()
const router = useRouter();

const username = ref('')
const password = ref('')

function handleSubmit() {
}

async function handleLogin() {
  try {
    await api_client.login(username.value, password.value)
  } catch (err) {
    if (err instanceof HTTPError) {
      toast(err.data.message)
      return
    }
    throw err
  }
  router.push({ name: 'home' })
}

async function handleRegister() {
  try {
    await api_client.register(username.value, password.value)
  } catch (err) {
    if (err instanceof HTTPError) {
      toast(err.data.message)
      return
    }
    throw err
  }
  await handleLogin()
}

</script>

<template>
  <form @submit.prevent="handleSubmit" class="max-w-md mx-auto p-6 bg-white shadow-lg rounded-lg">
    <h1 class="text-2xl font-semibold mb-6 text-center">Login</h1>
    <div class="mb-4">
      <label for="username" class="block text-sm font-medium text-gray-700 mb-2">Username:</label>
      <input type="text" id="username" v-model="username" required
        class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
    </div>
    <div class="mb-6">
      <label for="password" class="block text-sm font-medium text-gray-700 mb-2">Password:</label>
      <input type="password" id="password" v-model="password" required
        class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500">
    </div>
    <div class="flex justify-between">
      <button type="submit" @click="handleLogin"
        class="bg-blue-500 text-white py-2 px-4 rounded-lg hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500">Login</button>
      <button type="submit" @click="handleRegister"
        class="bg-green-500 text-white py-2 px-4 rounded-lg hover:bg-green-600 focus:outline-none focus:ring-2 focus:ring-green-500">Register</button>
    </div>
  </form>

</template>
