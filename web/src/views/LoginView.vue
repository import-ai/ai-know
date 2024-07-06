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
  <form @submit.prevent="handleSubmit">
    <h1>Login</h1>
    <div>
      <label for="username">Username:</label>
      <input type="text" id="username" v-model="username" required>
    </div>
    <div>
      <label for="password">Password:</label>
      <input type="password" id="password" v-model="password" required>
    </div>
    <div>
      <button type="submit" @click="handleLogin">Login</button>
      <button type="submit" @click="handleRegister">Register</button>
    </div>
  </form>
</template>
