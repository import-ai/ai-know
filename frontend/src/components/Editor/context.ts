import { AffineEditorContainer } from '@blocksuite/presets'
import { DocCollection } from '@blocksuite/store'
import { createContext, useContext } from 'react'

export const EditorContext = createContext<{
  editor: AffineEditorContainer
  collection: DocCollection
  renderMarkdown: (id: string, title: string, markdown: string) => Promise<void>
} | null>(null)

export function useEditor() {
  return useContext(EditorContext)
}
