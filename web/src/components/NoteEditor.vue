<script setup>
const props = defineProps(['kb'])

import { ref, watchEffect } from 'vue';
import { useToast } from "vue-toastification";

import { api_client } from '@/api_client'

const toast = useToast()

const notes = ref([])
const curNote = ref(null)
const editorTitle = ref('')
const editorContent = ref('')
const edited = ref(false)

watchEffect(async () => {
  if (props.kb) {
    notes.value = await api_client.listNotes(props.kb.id)
    if (notes.value.length > 0) {
      curNote.value = notes.value[0]
    } else {
      await createNewNote()
    }
  } else {
    notes.value = []
    curNote.value = null
  }
})

watchEffect(() => {
  if (curNote.value) {
    editorTitle.value = curNote.value.title
    editorContent.value = curNote.value.content
  } else {
    editorTitle.value = ''
    editorContent.value = ''
  }
  edited.value = false
})

async function createNewNote() {
  if (edited.value) {
    toast('Save first')
    return
  }
  const newNote = await api_client.createNote(props.kb.id, {})
  notes.value = await api_client.listNotes(props.kb.id)
  for (const note of notes.value) {
    if (note.id == newNote.id) {
      curNote.value = note
      return
    }
  }
  curNote.value = null
}

function truncate(content) {
  return content.length > 20 ? content.substring(0, 20) + '...' : content
}

function beautifyTitle(title) {
  if (title == "") {
    return '[No Title]'
  }
  return truncate(title)
}

function handleNoteSelect(note) {
  if (edited.value) {
    toast('Save first')
    return
  }
  curNote.value = note
}

async function handleInput() {
  edited.value = true
}

async function saveCurNote() {
  if (!curNote.value) {
    return
  }
  if (edited.value) {
    curNote.value.title = editorTitle.value
    curNote.value.content = editorContent.value
    await api_client.updateNote(props.kb.id, curNote.value)
    edited.value = false
    toast('Saved')
  }
}

</script>

<template>
  <div v-if="props.kb" class="flex size-full">
    <div class="w-1/4 flex flex-col overflow-y-auto">
      <button @click="createNewNote"
        class="w-full py-2 bg-green-500 text-white font-semibold rounded-lg shadow-md hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-400 focus:ring-opacity-75 mb-4">
        Create New Note
      </button>
      <div v-if="notes.length === 0" class="flex-1 flex flex-col justify-center">
        <p class="text-gray-500 text-center">No notes yet</p>
      </div>
      <ul v-else class="flex-1">
        <li v-for="note in notes" :key="note.id" :class="{
          'p-2 m-2 border border-gray-300 cursor-pointer': true,
          'bg-blue-100': curNote && note.id === curNote.id
        }" @click="handleNoteSelect(note)">
          <h3 class="font-bold">{{ beautifyTitle(note.title) }}</h3>
          <p>{{ truncate(note.content) }}</p>
        </li>
      </ul>
    </div>
    <div class="flex-1 ml-4 flex flex-col">
      <input class="w-full my-1 p-2 text-xl border border-gray-300 rounded" v-model="editorTitle" @input="handleInput"
        placeholder="Title" />
      <textarea class="w-full h-full p-2 border border-gray-300 rounded" v-model="editorContent" @input="handleInput"
        placeholder="Start typing..." />
      <button @click="saveCurNote"
        class="w-full my-1 py-2 bg-green-500 text-white font-semibold rounded-lg shadow-md hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-400 focus:ring-opacity-75">
        Save
      </button>
    </div>
  </div>
</template>
