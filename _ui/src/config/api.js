import axios from 'axios'

const apiConfig = {
  baseURL: '/api',
}

export const API = axios.create(apiConfig)

export const configForm = {
  headers: {
    'Content-Type': 'multipart/form-data',
  },
}

export const configJSON = {
  headers: {
    'Content-Type': 'application/json',
  },
}
