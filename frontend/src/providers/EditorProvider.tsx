import { EditorContext } from '@/components/Editor/context'
import { initEditor } from '@/components/Editor/editor'
import React from 'react'

export const EditorProvider = ({ children }: { children: React.ReactNode }) => {
  const { editor, collection } = initEditor()

  return (
    <EditorContext.Provider value={{ editor, collection }}>
      {children}
    </EditorContext.Provider>
  )
}
