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
    axios.get("http://127.0.0.1:3000/")
        .then(res => alert(res))
        .catch(err => alert(err))
  }
};
</script>

<style></style>