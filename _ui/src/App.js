import { BrowserRouter, Routes, Route } from 'react-router-dom'
import { UserProvider } from './contexts/UserContext'

import PrivatePage from './pages/PrivateRoute'
import Home from './pages/Home/Home'
import Admin from './pages/Admin/Admin'

export default function App() {
  return (
    <UserProvider>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route
            path="/admin"
            element={
              <PrivatePage>
                <Admin />
              </PrivatePage>
            }
          />
        </Routes>
      </BrowserRouter>
    </UserProvider>
  )
}
