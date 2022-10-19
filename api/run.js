const fetch = require('node-fetch')

export default async function run(request, response) {
  const resp = await fetch('https://go.dev/_/compile?backend=', {
    method: 'POST',
    headers: {
      'content-type': 'application/x-www-form-urlencoded'
    },
    body: new URLSearchParams({
      body: request.body.code,
      version: 2,
      withVet: true
    })
  })
  const data = await resp.json()
  response.status(200).json({
    err: data.Errors,
    events: data.Events
  })
}
