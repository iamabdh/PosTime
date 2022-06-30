<template>
  <div class="main-postime-create">
    <div class="postime-create">
      <form
          v-on:submit.prevent="sendPost"
      >
        <div class="postime-create-text">
        <textarea v-model="post" placeholder="Type here ..." required/>
        </div>
        <div class="postime-create-post-btn">
          <input type="submit" value="Post">
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "PosTimeCreate",
  data() {
    return {
      post: null
    }
  },
  methods: {
    sendPost() {
      axios.post("http://127.0.0.1:3000/user/postime/create", {
        text: this.post
      }, {
        withCredentials: true,
        headers: {
          'Content-Type' : 'application/x-www-form-urlencoded; charset=UTF-8'
        }
      })
          .then(res => console.log(res))
          .catch(err => console.log(err))
    },
  }
}
</script>

<style scoped>

.main-postime-create {
  /*border: 1px black solid;*/
  overflow: hidden;
  font-family: 'Poppins', sans-serif;
}
.postime-create-text {
  margin: 5px;
  text-align: center;
  padding: 10px;
}

textarea {
  resize: none;
  font-size: 17px;
  border: 1px #97ea7a solid;
  height: 100px;
  width: 100%;
  padding: 10px 0 0 10px;
  outline: none;
}

.postime-create-post-btn {
}
.postime-create-post-btn input {
  float: right;
  height: 48px;
  width: 166px;
  border: none;
  font-size: 20px;
  background: #97ea7a;
  border-radius: 20px;
  color: #000000;
}
.postime-create-post-btn input:hover {
  background: #b8ff93;
}
</style>