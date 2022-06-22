import React, { useEffect, useState } from 'react'
import { sortableContainer } from 'react-sortable-hoc'
import { arrayMoveImmutable } from 'array-move'
import { API, configJSON } from '../../../config/api'
import { useUser } from '../../../contexts/UserContext'
import Loading from '../../../components/Loading/Loading'
import './LinkEditor.css'
import LinkItem from './LinkItem'

const SortableContainer = sortableContainer(({ children }) => {
  return <div className="w-60">{children}</div>
})

export default function LinkEditor({
  links,
  setLinks,
  isFetchLoading,
  setFetchLoading,
}) {
  const { loggedInUser } = useUser()

  const fetchLinks = async () => {
    setFetchLoading(true)

    try {
      const response = await API.get(`/user/${loggedInUser.username}`, {
        withCredentials: true,
      })
      setLinks(response.data.data.links)
      setTimeout(() => setFetchLoading(false), 100)
    } catch (err) {
      console.log('error', `Error Fetch Data From API`)
    }
  }

  useEffect(() => {
    fetchLinks()
  }, [])

  const onSortEnd = ({ oldIndex, newIndex }) => {
    let newLinks
    setLinks((prev) => {
      let updatedLinks = arrayMoveImmutable(prev, oldIndex, newIndex)
      newLinks = updatedLinks
      return updatedLinks
    })
    updateLinksOrder(newLinks)
  }

  const updateLinksOrder = async (newLinks) => {
    let updatedLinks = newLinks.map((link) => link.id)
    try {
      const response = await API.patch(
        '/link/order',
        {
          links: updatedLinks,
        },
        {
          withCredentials: true,
          ...configJSON,
        }
      )
      console.log('success update order')
    } catch (err) {
      console.log(err)
    }
  }

  const deleteLink = async (linkId) => {
    setLinks((prev) => {
      return prev.filter((link) => link.id != linkId)
    })

    try {
      const response = await API.delete(`/link/${linkId}`, {
        withCredentials: true,
      })
      console.log('success delete link')
    } catch (err) {
      console.log(err)
    }
  }

  const createLink = async () => {
    try {
      const response = await API.post(`/link`, {
        withCredentials: true,
      })
      setLinks((prev) => {
        return [response.data.data, ...prev]
      })
      console.log('success create link')
    } catch (err) {
      console.log(err)
    }
  }

  return (
    <div>
      {isFetchLoading ? (
        <div className="w-100 d-flex jc-center">
          <Loading size="medium" color="gray" />
        </div>
      ) : (
        <div className="is-flex is-align-items-center is-flex-direction-column">
          <button class="button is-primary mb-5 w-30" onClick={createLink}>
            Add New Link
          </button>
          <SortableContainer onSortEnd={onSortEnd} useDragHandle>
            {links.map((link, index) => (
              <LinkItem
                key={`item-${link.id}`}
                index={index}
                value={link}
                deleteLink={deleteLink}
                setLinks={setLinks}
              />
            ))}
          </SortableContainer>
        </div>
      )}
    </div>
  )
}
