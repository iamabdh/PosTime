<template>
  <router-view />
</template>
<script>
import axios from "axios";

export default {
  name: "App",
  components: {},
  created() {
        switch (window.location.pathname) {
          case "/login":
            axios.get("http://127.0.0.1:3000/user/check/login", {
              withCredentials: true,
            })
                .then(res => {
                 this.$router.push(res.data.forward)
                })
                .catch(err => console.log(err))
                break
          case "/register":
            axios.get("http://127.0.0.1:3000/user/check/register", {
              withCredentials: true
            })
                .then(res => {
                  this.$router.push(res.data.forward)
                })
                .catch(err => console.log(err))
                break
          case "/":
            axios.get("http://127.0.0.1:3000/user/check/page", {
              withCredentials: true
            })
                .then(res => {
                  this.$router.push(res.data.forward)
                })
                .catch(err => console.log(err))
                break
        }
    },
};
</script>

<style></style>