<template>
  <div class="main-low-profile">
    <div class="image-low-profile">
      <img src="https://avatars.githubusercontent.com/u/20264401">
    </div>
    <div class="username-user-details-low-profile">
      <div class="username">
        <span v-if="username">@{{username}}</span>
      </div>
      <div class="name">
        {{nameProfile}}
      </div>
    </div>
    <div class="postimes-postimers-details-low-profile">
      <div class="postimes">
        PosTime <span style="margin: 0 20px">{{postime}}</span>
      </div>
      <div class="postimers">
        PosTimer <span style="margin: 0 20px">{{postimer}}</span>
      </div>
    </div>
    <div class="main-linkers-low-profile">
      <div>
        <button @click="changeComponentsHome" class="low-profile-btn">Home</button>
      </div>
      <div>
        <button @click="changeComponentsFindPosTimer" class="low-profile-btn">Find PosTimer</button>
      </div>
      <div>
        <button class="low-profile-btn">Settings</button>
      </div>
    </div>
    <div class="logout-low-profile">
      <button class="low-profile-btn">Logout</button>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "LowProfileHome",
  created() {
    this.getLowProfile()
  },
  data(){
    return{
      username: null,
      nameProfile: null,
      postime: null,
      postimer: null,
      homeFeedComponent: "PostimeViewHandlerComponent"
    }
  },
  methods: {
    getLowProfile() {
      axios.get("http://127.0.0.1:3000/user/postime/low-profile", {
        withCredentials: true
      })
          .then(res => {
            let lowProfileData = res.data
            this.username = lowProfileData.username
            this.nameProfile = lowProfileData.name
            this.postime = lowProfileData.postime
            this.postimer = lowProfileData.postimer
           })
          .catch(err => console.log(err))
    },
    changeComponentsHome() {
      this.homeFeedComponent = "PostimeViewHandlerComponent"
      this.$emit("Change-Component-Feed", this.homeFeedComponent)
    },
    changeComponentsFindPosTimer() {
      this.homeFeedComponent = "FindPosTimerHome"
      this.$emit("Change-Component-Feed", this.homeFeedComponent)
    }
  }
}
</script>

<style scoped>
* {
  font-family: 'Poppins', sans-serif;
}
.main-low-profile {
  text-align: center;
  margin: 10px 0;
}
.image-low-profile img {
  border-radius: 50%;
  width: 80px;
}
.username-user-details-low-profile > div,
.postimes-postimers-details-low-profile > div {
  margin: 10px 0;
}

.main-linkers-low-profile {
  margin: 100px 0;
}
.main-linkers-low-profile > div {
  margin: 10px 0;
}
.low-profile-btn {
  height: 30px;
  width: 184px;
  border: none;
  font-size: 16px;
  background: #12ff46;
  border-radius: 20px;
  color: #000000;
}
.logout-low-profile {
  position: relative;
  left: 50%;
  bottom: 10px;
  transform: translate(-50%, -50%);
  margin: 0 auto;
}
</style>