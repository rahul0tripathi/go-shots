const fetch = require('node-fetch')

export default async function fmt(request, response) {
  const resp = await fetch('https://go.dev/_/fmt?backend=', {
    method: 'POST',
    headers: {
      'content-type': 'application/x-www-form-urlencoded'
    },
    body: {
      body: request.body.code,
      imports: true
    }
  })
  console.log(resp, resp.json())
  const data = resp.json()
  response.status(200).json({
    err: data.Error,
    body: data.Body
  })
}
