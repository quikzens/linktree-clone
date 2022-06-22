import LinkEditor from './LinkEditor/LinkEditor'
import { useState } from 'react'
import LinkPreview from './LinkPreview/LinkPreview'

export default function Admin() {
  const [links, setLinks] = useState([])
  const [isFetchLoading, setFetchLoading] = useState(true)

  return (
    <div className="columns">
      <div className="column p-6">
        <LinkEditor
          links={links}
          setLinks={setLinks}
          isFetchLoading={isFetchLoading}
          setFetchLoading={setFetchLoading}
        />
      </div>
      <div className="column is-one-third p-6">
        <LinkPreview
          links={links}
          isFetchLoading={isFetchLoading}
          setFetchLoading={setFetchLoading}
        />
      </div>
    </div>
  )
}
