import { API } from '../config/api'

export const fetchData = async (url, setData, setLoading) => {
  setLoading(true)

  try {
    const response = await API.get(url, { withCredentials: true })
    setData(response.data.data)
    setTimeout(() => setLoading(false), 100)
  } catch (err) {
    console.log('error', `Error Fetch Data From API`)
  }
}
