import { Link } from 'react-router-dom'
import { useUser } from '../../contexts/UserContext'
import './Home.css'

export default function Home() {
  const { isLoggedIn } = useUser()

  return (
    <div className="page-home is-flex is-justify-content-center is-align-items-center">
      <div className="home-container w-60 is-flex is-flex-direction-column is-align-items-center">
        <h1 className="title">
          Linktree <br /> Clone
        </h1>

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
  )
}
