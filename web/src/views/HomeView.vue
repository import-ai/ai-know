<script setup>
import { ref, watchEffect } from 'vue'
import { useRouter } from 'vue-router';

import { HTTPError } from '@/http_client'
import { api_client } from '@/api_client'

const router = useRouter();

const authorized_user = ref('')
watchEffect(async () => {
  try {
    authorized_user.value = await api_client.getAuthorizedUser()
  } catch (err) {
    if (err instanceof HTTPError && err.status == 401) {
      router.push('login')
      return
    }
    throw err
  }
})

</script>

<template>
  <main>
    <h1>Home</h1>
    <p>Logged in as: {{ authorized_user }}</p>
  </main>
</template>
