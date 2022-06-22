import { useState, useEffect } from 'react'
import { useParams } from 'react-router-dom'
import { API } from '../../config/api'
import Loading from '../../components/Loading/Loading'

export default function Links() {
  let { username } = useParams()
  const [user, setUser] = useState({})
  const [isFetchLoading, setFetchLoading] = useState(true)

  const fetchUser = async () => {
    setFetchLoading(true)

    try {
      const response = await API.get(`/user/${username}`, {
        withCredentials: true,
      })
      setUser(response.data.data)
      setTimeout(() => setFetchLoading(false), 100)
    } catch (err) {
      console.log('error', `Error Fetch Data From API`)
    }
  }

  useEffect(() => {
    fetchUser()
  }, [])

  return (
    <div className="link-preview">
      {isFetchLoading ? (
        <div className="w-100 d-flex jc-center">
          <Loading size="medium" color="gray" />
        </div>
      ) : (
        <>
          <img src={user.avatar_url} className="link-preview-avatar" />
          <h1 className="link-preview-username">@{user.username}</h1>
          <div className="link-preview-list">
            {user.links.map((link) => {
              if (link.is_active) {
                return (
                  <a className="link-preview-item" href={link.url}>
                    {link.title}
                  </a>
                )
              }
            })}
          </div>
        </>
      )}
    </div>
  )
}
