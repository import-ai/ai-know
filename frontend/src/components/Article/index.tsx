import EditorContainer from '@/components/Editor/EditorContainer'
import { useParams } from 'react-router-dom'

export const Article = () => {
  const { id } = useParams()

  return (
    <div className="p-10 bg-slate-50">
      <code>Article {id}</code>
      <EditorContainer />
    </div>
  )
}
