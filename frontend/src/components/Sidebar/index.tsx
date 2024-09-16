import { useArticleList } from '@/hooks/article'
import { Link } from 'react-router-dom'
import { useEffect, useRef, useState } from 'react'
import Sortable from 'sortablejs'

interface Article {
  id: string
  title: string
  content: string
}

export const Sidebar = () => {
  const list = useArticleList()
  const sortableRef = useRef<HTMLUListElement>(null)
  const [selectedId, setSelectedId] = useState<string | null>(null)

  useEffect(() => {
    if (sortableRef.current) {
      Sortable.create(sortableRef.current, {
        animation: 150,
        onEnd: (evt) => {
          const { oldIndex, newIndex } = evt
          if (oldIndex !== newIndex) {
            console.log(`Moved item from ${oldIndex} to ${newIndex}`)
          }
        },
      })
    }
  }, [])

  return (
    <div className="drawer-open">
      <input id="my-drawer-2" type="checkbox" className="drawer-toggle" />

      <div className="drawer-side">
        <label
          htmlFor="my-drawer-2"
          aria-label="close sidebar"
          className="drawer-overlay"
        ></label>
        <ul
          className="menu bg-base-200 text-base-content min-h-full w-80 p-4"
          ref={sortableRef}
        >
          {list.map((article: Article) => (
            <li key={article.id}>
              <Link
                to={`/article/${article.id}`}
                className={`hover:bg-base-300 ${
                  selectedId === article.id ? 'bg-base-300 font-bold' : ''
                }`}
                onClick={() => setSelectedId(article.id)}
              >
                {article.title}
              </Link>
            </li>
          ))}
          <div className="flex-1"></div>
          <div>
            {/* debug tool */}
            <button className="btn">export snap</button>
            <button className="btn">console markdown</button>
          </div>
        </ul>
      </div>
    </div>
  )
}