<template>
  <div class="main-postimer-update">
    <div class="title-postimer-update">
      <h4>
        PosTimers Update
      </h4>
    </div>
    <div class="users-postimer-update-main">
      <div class="container-main-user-detail-postimer">
        <div v-for="data in lastUpdateData">
          <div class="user-postimer" :aria-label="data.PosTimeId" @click="this.updateUserScroll(data.PosTimeID)">
            <div class="user-image-postimer">
              <img src="https://avatars.githubusercontent.com/u/20264401">
            </div>
            <div class="user-details-postimer">
              <div class="time-latest-postimer">
                {{getElapsedTime(data.date)}}
              </div>
              <div class="username-postimer">
                {{data.username}}
              </div>
            </div>
          </div>
        </div>
      </div>
      </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "PostimersUpdateHome",
  // emits: {
  //   scroll: () => {
  //     alert("ss")
  //   }
  // },
  data(){
    return {
      lastUpdateData: null,
    }
  },
  created() {
     this.getLastUpdate()
  },
  methods: {
    getLastUpdate() {
      axios.get("http://127.0.0.1:3000/user/postime/last-update", {
        withCredentials: true
      })
          .then(res => {
            this.lastUpdateData = res.data
          })
          .catch(err => console.log(err))
    },
    updateUserScroll(id) {
    },
    getElapsedTime (t1) {
      let seconds =Math.floor((Date.now() -Date.parse(t1)) /1000)
      if (seconds < 60) {
        return seconds + " s"
      }
      let mins = Math.floor(seconds/60)
      if (mins < 60) {
        return mins + " m"
      }
      let hours =Math.floor(mins/60)
      if (hours < 24) {
        return hours + " h"
      }
      let days = Math.floor(hours/24)
      return days + " d"
    }
  }
}
</script>

<style scoped>
.main-postimer-update {
  text-align: center;
}
.users-postimer-update-main {
  margin: 30px 0;
  font-size: 12px;
}
.container-main-user-detail-postimer {
  margin: 10px 0;
  /*border: 1px solid;*/
}
.user-postimer {
  display: flex;
  justify-content: center;
  margin: 0 auto;
  /*border: 1px black solid;*/
  width: 200px;
  /*align-items: center;*/
  padding-top: 5px;

}
.user-postimer:hover {
  background: #d1e1d5;
}
.user-postimer > div {
  margin: 0 2px;
}
.user-image-postimer img{
  border-radius: 50%;
  width: 42px;
}
.user-details-postimer {
  padding-top: 3px;
}
.user-details-postimer .time-latest-postimer {
  float: left;
}
</style>