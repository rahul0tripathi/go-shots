import axios from 'axios'
var base64 = require('js-base64').Base64

const repoURI = `rahul0tripathi/go-shots`

export const getAllFilesInRepo = async () => {
  const response = await axios.get(
    ` https://api.github.com/repos/${repoURI}/git/trees/master?recursive=1`
  )
  const files = response.data.tree.map((v) => {
    return {
      fileName: v.path,
      resourceURI: v.url
    }
  })
  console.log(files)
  return files
}

export const fetchContent = async (uri) => {
  const response = await axios.get(`${uri}`)
  const content = base64.decode(response.data.content)
  return content
}
