import Loading from '../../../components/Loading/Loading'
import { useUser } from '../../../contexts/UserContext'
import './LinkPreview.css'

export default function LinkPreview({ links, isFetchLoading }) {
  const { loggedInUser } = useUser()

  return (
    <div className="link-preview mobile-preview">
      {isFetchLoading ? (
        <div className="w-100 d-flex jc-center">
          <Loading size="medium" color="gray" />
        </div>
      ) : (
        <>
          <img
            src={loggedInUser.user_avatar_url}
            className="link-preview-avatar"
          />
          <h1 className="link-preview-username">@{loggedInUser.username}</h1>
          <div className="link-preview-list">
            {links.map((link) => {
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
