import React, { useState } from 'react'
import { sortableElement, sortableHandle } from 'react-sortable-hoc'
import { BsThreeDotsVertical, BsTrash } from 'react-icons/bs'
import './LinkEditor.css'
import Switch from '@mui/material/Switch'
import { API, configJSON } from '../../../config/api'

const DragHandle = sortableHandle(() => (
  <span>
    <BsThreeDotsVertical />
  </span>
))

const LinkItem = sortableElement(({ value, deleteLink }) => {
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
    console.log(name, value)
    setLink((prev) => {
      return {
        ...prev,
        [name]: value,
      }
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
    <div className="link-item is-flex card mb-4 is-two-thirds w-60">
      <div className="p-4 is-flex is-align-items-center">
        <DragHandle />
      </div>
      <div className="is-flex is-justify-content-space-between p-4 w-100">
        <div>
          <form className="is-flex is-flex-direction-column">
            <input
              type="text"
              defaultValue={link.title}
              onBlur={(e) => updateLink('title', e.target.value)}
              onKeyDown={blurOnEnter}
            />
            <input
              type="text"
              defaultValue={link.url}
              onBlur={(e) => updateLink('url', e.target.value)}
              onKeyDown={blurOnEnter}
            />
          </form>
        </div>
        <div className="is-flex is-flex-direction-column is-align-items-end">
          <Switch
            checked={link.is_active}
            onClick={() => updateLink('is_active', !link.is_active)}
          />
          <div onClick={() => deleteLink(link.id)}>
            <BsTrash />
          </div>
        </div>
      </div>
    </div>
  )
})

export default LinkItem
