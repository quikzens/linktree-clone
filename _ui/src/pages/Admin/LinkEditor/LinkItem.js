import React, { useState } from 'react'
import { sortableElement, sortableHandle } from 'react-sortable-hoc'
import { BsThreeDotsVertical, BsTrash } from 'react-icons/bs'
import './LinkEditor.css'
import Switch from '@mui/material/Switch'
import { API, configJSON } from '../../../config/api'

const DragHandle = sortableHandle(() => (
  <div className="link-item-dragger p-4 is-flex is-align-items-center">
    <span>
      <BsThreeDotsVertical />
    </span>
  </div>
))

const LinkItem = sortableElement(({ value, deleteLink, setLinks }) => {
  const [link, setLink] = useState({
    id: value.id,
    title: value.title,
    url: value.url,
    is_active: value.is_active,
  })

  const blurOnEnter = (e) => {
    if (e.code === 'Enter') {
      e.target.blur()
    }
  }

  const updateLink = async (name, value) => {
    setLink((prev) => {
      return {
        ...prev,
        [name]: value,
      }
    })
    setLinks((prev) => {
      return prev.map((item) => {
        if (link.id === item.id) {
          return {
            ...item,
            [name]: value,
          }
        }
        return item
      })
    })

    try {
      const response = await API.patch(
        `/link/${link.id}`,
        {
          [name]: value,
        },
        {
          withCredentials: true,
          ...configJSON,
        }
      )
      console.log('success update link')
    } catch (err) {
      console.log(err)
    }
  }

  return (
    <div className="link-item is-flex card mb-4">
      <DragHandle />
      <div className="is-flex is-justify-content-space-between is-align-items-center p-4 w-100">
        <div>
          <form className="is-flex is-flex-direction-column">
            <input
              type="text"
              className="link-item-input input-title"
              defaultValue={link.title}
              onBlur={(e) => updateLink('title', e.target.value)}
              onKeyDown={blurOnEnter}
            />
            <input
              type="text"
              className="link-item-input input-url"
              defaultValue={link.url}
              onBlur={(e) => updateLink('url', e.target.value)}
              onKeyDown={blurOnEnter}
            />
          </form>
        </div>
        <div className="link-item-cta is-flex is-flex-direction-column is-justify-content-space-between is-align-items-end">
          <Switch
            checked={link.is_active}
            onClick={() => updateLink('is_active', !link.is_active)}
          />
          <div onClick={() => deleteLink(link.id)} className="link-item-del">
            <BsTrash />
          </div>
        </div>
      </div>
    </div>
  )
})

export default LinkItem
