const fetch = require('node-fetch')

export default async function fmt(request, response) {
  const resp = await fetch('https://go.dev/_/fmt?backend=', {
    method: 'POST',
    headers: {
      'content-type': 'application/x-www-form-urlencoded'
    },
    body: new URLSearchParams({
      body: request.body.code,
      imports: true
    })
  })

  const data = resp.json()
  console.log(resp, data)
  response.status(200).json({
    err: data.Error,
    body: data.Body
  })
}
