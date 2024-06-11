'use client'
import { useState, useEffect } from 'react';

import NoteList from '@/components/NoteList'
import NoteEditor from '@/components/NoteEditor'

interface Note {
  id: number,
  content: string,
}

const getAllNotes = async () => {
  const resp = await fetch('/api/get_all_notes', {
    method: 'POST',
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({})
  })
  const notes = await resp.json()
  return notes['notes'].map((note) => (
    { id: note['id'], content: note['content'] }
  ))
}

const updateNote = async (id: number, newContent: string) => {
  const resp = await fetch('/api/update_note', {
    method: 'POST',
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ id: id, content: newContent })
  })
}

const createNote = async (content: string) => {
  const resp = await fetch('/api/create_note', {
    method: 'POST',
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ content: content })
  })
  const respJson = await resp.json()
  return respJson.id
}

const getNoteByID = async (id: number) => {
  const resp = await fetch('/api/get_note_by_id', {
    method: 'POST',
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ id: id })
  })
  const respJson = await resp.json()
  console.log(respJson)
  return respJson.note

}

const Home = () => {
  const [needRefresh, setNeedRefresh] = useState<boolean>(true)
  const [notes, setNotes] = useState<Note[]>([])
  const [selectedNote, setSelectedNote] = useState<Note>({ id: -1, content: '' })

  if (needRefresh) {
    const f = async () => {
      const allNotes = await getAllNotes()
      setNotes(allNotes)
    }
    setNeedRefresh(false)
    f().catch()
  }

  async function onNoteSelect(id: number) {
    if (selectedNote.id != -1 || selectedNote.content != '') {
      if (selectedNote.id == -1) {
        await createNote(selectedNote.content)
      } else {
        updateNote(selectedNote.id, selectedNote.content)
      }
    }
    let newContent = ''
    if (id != -1) {
      const note = await getNoteByID(id)
      newContent = note.content
    }
    setSelectedNote({ id: id, content: newContent })
  }

  async function onNoteUpdate(newContent: string) {
    let id = selectedNote.id
    if (id == -1) {
      id = await createNote(newContent)
    } else {
      updateNote(id, newContent)
    }
    setNeedRefresh(true)
    setSelectedNote({ id: id, content: newContent })
  }
  console.log(notes)

  return (
    <main className="flex h-screen">
      <NoteList notes={notes} onSelect={onNoteSelect}></NoteList>
      <NoteEditor content={selectedNote.content} onUpdate={onNoteUpdate}></NoteEditor>
    </main>
  );
}

export default Home