<template>
  <div id="main-container">
    <div id="form-container">
      <form
      @submit="sendData()"
      >
        <div class="form-group">
          <label for="name">Name</label>
          <input type="text" v-model="name" />
        </div>
        <div class="form-group">
          <label for="email">Email</label>
          <input type="text" v-model="email" />
        </div>
        <div class="form-group">
          <label for="username">Username</label>
          <input type="text" v-model="username" />
        </div>
        <div class="form-group">
          <label for="password">Password</label>
          <input type="text" v-model="password" />
        </div>
        <div class="form-group">
          <label for="confirmPassword">Confirm Password</label>
          <input type="text" v-model="confirmPassword" />
        </div>
        <div id="status-notifiers">
        </div>
        <input type="submit" value="Sign Up">
      </form>
    </div>
<!--    <div id="brand-container">-->
<!--    <div id="brand-title">-->
<!--      <h2>-->
<!--        PosTime-->
<!--      </h2>-->
<!--    </div>-->
<!--    </div>-->
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      name: null,
      email: null,
      username: null,
      password: null,
      confirmPassword: null,
    }
  },
  methods: {
    sendData(){
      // validate before request
      axios
          .post("http://127.0.0.1:3000/user/register", {
            Name: this.name,
            Email: this.email,
            Username: this.username,
            Password: this.password
          })
          .then(res => {console.log(res)})
          .catch(err => {console.log(err)})
    }
  }
};
</script>
<style scoped>
* {
  margin: 0;
  padding: 0;
  overflow: hidden;
}
#main-container {
  display: flex;
}
#form-container {
  /*border: black 1px solid;*/
  text-align: center;
  width: 60%;
}

#brand-container {
  width: 40%;
  background: #97ea7a;
  position: absolute;
  top: 0;
  right: 0;
  bottom: 0;
}

#brand-title {
  line-height: 200px;
  height: 100%;
  /*border: 3px solid green;*/
  text-align: center;
}

#brand-title p {
  vertical-align: middle;
  display: inline-block;
  line-height: 1.5;
}
label {
  font-family: Georgia, "Times New Roman", Times, serif;
  font-size: 18px;
  color: #333;
  height: 20px;
  width: 200px;
  margin-top: 10px;
  text-align: left;
  /*margin-left: 10px;*/
  /*text-align: right;*/
  /*margin-right:15px;*/
  float: left;
}
input {
  height: 20px;
  width: 200px;
  border: 1px solid #000;
  margin-top: 10px;
}
</style>