export default function ({ $axios, redirect }) {
  // $axios.onRequest((config) => {})

  $axios.onResponse((response) => {
    return Promise.resolve(response.data || response)
  })

  $axios.onError((error) => {
    try {
      const code = parseInt(error.response && error.response.status)
      if (code === 400) {
        redirect('/400')
      }
    } catch (error) {
      redirect('/500')
    }
  })
}
