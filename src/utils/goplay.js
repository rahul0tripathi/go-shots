import axios from 'axios'

export const fmt = async (text) => {
  const fmtData = await axios.post(`${window.location.origin}/api/fmt`, {
    code: text
  })
  return fmtData.data
}

export const run = async (text) => {
  const runData = await axios.post(`${window.location.origin}/api/run`, {
    code: text
  })
  return runData.data
}
