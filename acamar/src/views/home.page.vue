<template>
  <div id="main-container">
    <div>
    </div>
    <div>
      data user: <span>
      name: {{name}}
      email: {{email}}
      username:{{username}}
    </span>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Home",
  title: "Home",
  data() {
    return {
      name: null,
      email: null,
      username:null,
    }
  },
  created() {
    axios.get("http://127.0.0.1:3000/user/page", {
      withCredentials: true,
    })
        .then(res => {
          this.getData()
        })
        .catch(err => {
          this.$router.push(err.response.data.forward)
        })
  },
  methods: {
    getData() {
      axios.get("http://127.0.0.1:3000/user/page", {
        withCredentials: true,
      })
          .then(res => {
            let data = JSON.parse(res.data)
            this.name = data.name
            this.email = data.email
            this.username = data.username
          })
          .catch(err => console.log(err))
    }
  }
}
</script>

<style scoped>

</style>