export default ({ $axios, store }) => {
  $axios.onResponse((response) => {
    console.log(`[${response.status}] ${response.request.path}`)

    console.log('base url', process.env.API_URL)
  })

  $axios.onError((err) => {
    console.log(`[${err.response && err.response.status}] ${err.response && err.response.request.path}`)
    console.log(err.response && err.response.data)
  })
}
