import { Navigate } from 'react-router-dom'
import { useUser } from '../contexts/UserContext'

/**
 * Wrapper component to protect particular route
 * we use conditional rendering base on state to check
 * if there is user login or not
 * if login: render page (children)
 * if not: redirect to home path
 */
export default function PrivatePage(props) {
  const { isLoggedIn } = useUser()

  if (!isLoggedIn) {
    return <Navigate replace to={'/'} />
  }

  return props.children
}
