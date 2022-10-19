const axios = require('axios')

export default async function run(request, response) {
  const resp = await axios.post(
    'https://go.dev/_/compile?backend=',
    {
      body: request.body.code,
      version: 2,
      withVet: true
    },
    {
      headers: {
        'content-type': 'application/x-www-form-urlencoded'
      }
    }
  )
  response.status(200).json({
    err: resp.data.Errors,
    events: resp.data.Events
  })
}
