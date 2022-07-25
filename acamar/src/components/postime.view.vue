<template>
  <div class="main-postime-view">
    <div class="postime-view" :id="postime.PosTimeID">
      <div class="profile-user-postime-view">
        <img src="https://avatars.githubusercontent.com/u/20264401?s=64&v=4">
        <span class="profile-user-postime-view-username">@{{postime.Username}}</span>
        <span class="profile-user-postime-view-updated-time-date">{{getElapsedTime(postime.Date)}}</span>
      </div>
      <div class="postime-view-postime-content">
        <p>
          {{postime.Text}}
        </p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "PosTimeView",
  props: ['postime'],
  methods: {
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
    },
    scrollPosTime() {
      let emitter = require('tiny-emitter/instance');
      emitter.on('some-event', (id) => alert(id))
    }
  }
}
</script>

<style scoped>
  .main-postime-view {
    margin: 10px;
  }
  .postime-view {
    border: 1px #97ea7a solid;
  }
  .profile-user-postime-view {
    display: flex;
    margin: 4px;
  }
  .profile-user-postime-view img {
    border-radius: 50%;
    width: 55px;
    margin-right: 5px;
  }
  .profile-user-postime-view-username {
    margin-top: 15px;
  }
  .profile-user-postime-view-updated-time-date {
    margin-left: auto;
    margin-right: 30px;
  }
  .postime-view-postime-content {
    margin: 10px;
  }
</style>