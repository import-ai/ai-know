import { EditorContext } from '@/components/Editor/context'
import { initEditor } from '@/components/Editor/editor'
import React from 'react'

export const EditorProvider = ({ children }: { children: React.ReactNode }) => {
  const { editor, collection, renderMarkdown } = initEditor()

  return (
    <EditorContext.Provider value={{ editor, collection, renderMarkdown }}>
      {children}
    </EditorContext.Provider>
  )
}
