import { Link } from 'react-router-dom'
import { useUser } from '../../contexts/UserContext'

export default function Home() {
  const { isLoggedIn } = useUser()

  return (
    <div>
      <nav
        className="navbar is-flex is-justify-content-space-between"
        role="navigation"
        aria-label="main navigation"
      >
        <div className="navbar-brand">
          <a className="navbar-item is-size-4 has-text-weight-bold" href="/">
            Linktree Clone
          </a>
        </div>

        <div className="navbar-end">
          <div className="navbar-item">
            <div className="buttons">
              {isLoggedIn ? (
                <>
                  <Link to="/admin" className="button is-primary">
                    <strong>Admin</strong>
                  </Link>
                  <a className="button" href="/auth/logout">
                    <strong>Log Out</strong>
                  </a>
                </>
              ) : (
                <a className="button is-primary" href="/auth/login/google">
                  <strong>Log In</strong>
                </a>
              )}
            </div>
          </div>
        </div>
      </nav>
    </div>
  )
}
