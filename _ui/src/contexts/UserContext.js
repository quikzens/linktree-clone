import React, { createContext, useState, useEffect, useContext } from 'react'
import { API } from '../config/api'

export const UserContext = createContext()

export const UserProvider = ({ children }) => {
  const [loggedInUser, setLoggedInUser] = useState({})
  const [isLoggedIn, setLoggedIn] = useState(false)

  useEffect(() => {
    invokeAuth()
  }, [])

  const invokeAuth = async () => {
    try {
      const response = await API.get('/user/auth', { withCredentials: true })
      setLoggedIn(true)
      setLoggedInUser(response.data.data)
    } catch (err) {
      setLoggedIn(false)
      setLoggedInUser({})
    }
  }

  return (
    <UserContext.Provider
      value={{
        isLoggedIn,
        loggedInUser,
        setLoggedInUser,
        invokeAuth,
      }}
    >
      {children}
    </UserContext.Provider>
  )
}

export const useUser = () => {
  const context = useContext(UserContext)
  if (context === undefined) {
    throw new Error('useUser must be used within a UserProvider')
  }
  return context
}
