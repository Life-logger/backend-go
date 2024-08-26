import axios from 'axios'

const baseURL = 'http://localhost:22250/'

export default axios.create(
    {
        baseURL,
        headers: {"Content-Type": "application/json"},
        withCredentials: true,
    }
)
