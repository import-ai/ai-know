<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router';

import { HTTPError } from '@/http_client'
import { api_client } from '@/api_client'
import NavBar from '@/components/NavBar.vue'
import CreateKbModal from '@/components/CreateKbModal.vue'

const router = useRouter();

const authorizedUser = ref('')
const kbs = ref([])
const curKb = ref(null)
const isModalOpen = ref(false)

onMounted(async () => {
  try {
    authorizedUser.value = await api_client.getAuthorizedUser()
    await refreshKbs()
  } catch (err) {
    if (err instanceof HTTPError && err.status == 401) {
      router.push('login')
      return
    }
    throw err
  }
})

function isCurKbValid() {
  return curKb.value && kbs.value.some((kb) => kb.id == curKb.value.id)
}

async function refreshKbs() {
  kbs.value = await api_client.listKbs()
  if (!isCurKbValid() && kbs.value.length > 0) {
    curKb.value = kbs.value[0]
  }
}

async function handleLogout() {
  await api_client.logout()
  router.push('login')
}

async function handleKbSelected(kb) {
  curKb.value = kb
}

function handleCreateKbClicked() {
  isModalOpen.value = true
}

async function handleModalSubmit(newKbTitle) {
  curKb.value = await api_client.createKb({ title: newKbTitle })
  await refreshKbs()
  isModalOpen.value = false
}
function handleModalClose() {
  isModalOpen.value = false
}

</script>

<template>
  <NavBar :kbs="kbs" :cur-kb="curKb" :user-name="authorizedUser" @kb-selected="handleKbSelected"
    @logout-clicked="handleLogout" @create-kb-clicked="handleCreateKbClicked"></NavBar>
  <CreateKbModal :is-open="isModalOpen" @submit="handleModalSubmit" @close="handleModalClose"></CreateKbModal>
</template>
