import { Link } from 'react-router-dom'

export const Sidebar = () => {
  return (
    <div
      style={{ width: '200px', borderRight: '1px solid #ccc', padding: '10px' }}
    >
      <ul>
        <li>
          <Link to="/article/1">Article 1</Link>
        </li>
        <li>
          <Link to="/article/2">Article 2</Link>
        </li>
        <li>
          <Link to="/article/3">Article 3</Link>
        </li>
      </ul>
    </div>
  )
}
