<template>
  <router-view />
</template>
<script>
import axios from "axios";

export default {
  name: "App",
  components: {},
  created() {
  //  GET request to server
    axios.get("http://127.0.0.1:3000/", {
     withCredentials: true,
    })
        .then(res => console.log(res.data))
        .catch(err => console.log(err))
  }
};
</script>

<style></style>