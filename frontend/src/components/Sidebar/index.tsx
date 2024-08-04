import { useArticleList } from '@/hooks/article'
import { Link } from 'react-router-dom'

export const Sidebar = () => {
  const list = useArticleList()
  return (
    <div
      style={{ width: '200px', borderRight: '1px solid #ccc', padding: '10px' }}
    >
      <ul>
        {list.map((article) => (
          <li key={article.id}>
            <Link to={`/article/${article.id}`}>{article.title}</Link>
          </li>
        ))}
      </ul>
    </div>
  )
}
