import { useEditor } from '@/components/Editor/context'
import EditorContainer from '@/components/Editor/EditorContainer'
import { useArticleList } from '@/hooks/article'
import { useEffect, useMemo } from 'react'
import { useParams } from 'react-router-dom'

export const Article = () => {
  const { id } = useParams()
  const list = useArticleList()
  const content = useMemo(
    () => list.find((article) => article.id === id),
    [list, id],
  )
  const { renderMarkdown } = useEditor()!
  useEffect(() => {
    if (content) {
      renderMarkdown(content.id, content.title, content.content)
    }
  }, [content])
  return (
    <div className="p-10 bg-slate-50">
      <code>Article {id}</code>
      <EditorContainer />
    </div>
  )
}
