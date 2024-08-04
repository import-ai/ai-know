import { useParams } from 'react-router-dom'

export const Article = () => {
  const { id } = useParams()

  return (
    <div className="p-10 bg-slate-100">
      <h1>Article {id}</h1>
      <p>This is the content of article {id}.</p>
    </div>
  )
}
