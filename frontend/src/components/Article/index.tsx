import { useParams } from 'react-router-dom'

export const Article = () => {
  const { id } = useParams()

  return (
    <div style={{ padding: '10px' }}>
      <h1>Article {id}</h1>
      <p>This is the content of article {id}.</p>
    </div>
  )
}
