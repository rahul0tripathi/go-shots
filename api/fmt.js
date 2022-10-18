import axios from 'axios'

export default async function fmt(request, response) {
  const resp = await axios.post(
    'https://go.dev/_/fmt?backend=',
    {
      body: request.body.code,
      imports: true
    },
    {
      headers: {
        'content-type': 'application/x-www-form-urlencoded'
      }
    }
  )
  response.status(200).json({
    err: resp.data.Error,
    body: resp.data.Body
  })
}
