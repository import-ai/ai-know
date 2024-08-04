import { AffineEditorContainer } from '@blocksuite/presets'
import { Doc, nanoid, Schema } from '@blocksuite/store'
import { DocCollection } from '@blocksuite/store'
import { AffineSchemas, MarkdownTransformer } from '@blocksuite/blocks'
import '@blocksuite/presets/themes/affine.css'

export interface EditorContextType {
  editor: AffineEditorContainer | null
  collection: DocCollection | null
  updateCollection: (newCollection: DocCollection) => void
}

export function initEditor() {
  const schema = new Schema().register(AffineSchemas)
  const collection = new DocCollection({ schema })
  collection.meta.initialize()

  const doc = collection.createDoc({ id: 'page1' })
  doc.load(() => {
    const pageBlockId = doc.addBlock('affine:page', {})
    doc.addBlock('affine:surface', {}, pageBlockId)
    const noteId = doc.addBlock('affine:note', {}, pageBlockId)
    doc.addBlock('affine:paragraph', {}, noteId)
  })

  const editor = new AffineEditorContainer()
  editor.doc = doc
  // editor.slots.docLinkClicked.on(({ docId }) => {
  //   const target = <Doc>collection.getDoc(docId)
  //   editor.doc = target
  // })

  const renderMarkdown = async (
    id: string,
    title: string,
    markdown: string,
  ) => {
    const doc = collection.createDoc({ id: nanoid() })
    editor.doc = doc
    doc.load()
    // Add root block and surface block at root level
    const rootId = doc.addBlock('affine:page', {
      // title: new Text(title),
    })
    doc.addBlock('affine:surface', {}, rootId)

    const noteId = doc.addBlock(
      'affine:note',
      { xywh: '[0, 100, 800, 640]' },
      rootId,
    )

    await MarkdownTransformer.importMarkdown({
      doc,
      noteId,
      markdown: markdown,
    })

    doc.resetHistory()
  }
  return { editor, collection, renderMarkdown }
}
