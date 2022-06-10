<template>
  <form
  @submit="sendData()">
    <input type="text" placeholder="name" v-model = "name">
    <input type="text" placeholder="email" v-model = "email">
    <input type="text" placeholder="username" v-model = "username">
    <input type="password" placeholder="password" v-model = "password">
    <input type="submit" value="send">
  </form>
  <button @click="clickMe()">test</button>
</template>

<script>
import axios from "axios";
export default {
  name: "Register",
  title: "Register",
  data() {
    return {
      name: null,
      email: null,
      username: null,
      password: null,
    }
  },
  methods: {
    sendData(){
      axios
          .post("http://127.0.0.1:3000/user/register", {
            Name: this.name,
            Email: this.email,
            Username: this.username,
            Password: this.password
          })
          .then(res => {alert(res)})
          .catch(err => {alert(err)})
    },
    clickMe(){

    }
  },
}
</script>

<style scoped>

</style>