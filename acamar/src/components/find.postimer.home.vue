<template>
  <div class="main-find-postimer">
    <div class="container-card-find-postimer">
      <div v-for="objPosTimer in objNewPosTimer">
        <div class="card-find-postimer">
          <div class="image-container-find-postimer">
            <img src="https://avatars.githubusercontent.com/u/20264401" />
          </div>
          <div
            class="
              user-details-find-postimer
              verical-alingment-postimes-postimer-common
            "
          >
            <div class="name-find-postimer">
              {{ objPosTimer.name }}
            </div>
            <div class="username-find-postimer">
              @{{ objPosTimer.username }}
            </div>
          </div>
          <div
            class="
              postime-postimers-find-postimer
              verical-alingment-postimes-postimer-common
            "
          >
            <div class="postime-find-postimer">
              PosTime
              <span class="postime-postimers-values">{{
                objPosTimer.postime
              }}</span>
            </div>
            <div class="postimer-find-postimer">
              PosTimer
              <span class="postime-postimers-values">{{
                objPosTimer.postimer
              }}</span>
            </div>
          </div>
          <div
            class="
              last-update-follow-find-postimer
              verical-alingment-postimes-postimer-common
            "
            style="margin-top: 15px"
          >
            <div v-if="objPosTimer.lastUpdate">
              <div class="last-update">
                Last Update
                <span></span
                ><span class="last-update-value">{{
                  this.getElapsedTime(objPosTimer.lastUpdate)
                }}</span>
              </div>
            </div>
            <div class="follow-find-postimer">
              <button
                  class="follow-btn"
                :aria-label="objPosTimer.username"
                @click="this.addNewPosTimer(objPosTimer.username)"
              >
                Follow
              </button>
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
  name: "FindPosTimerHome",
  data() {
    return {
      objNewPosTimer: null,
    };
  },
  created() {
    this.getNewPosTimer();
  },
  methods: {
    getNewPosTimer() {
      axios
        .get("http://127.0.0.1:3000/user/postime/find-postimer", {
          withCredentials: true,
        })
        .then((res) => {
          this.objNewPosTimer = res.data;
        })
        .catch((err) => console.log(err));
    },
    addNewPosTimer(username) {
      // select element
      const element = document.querySelector(`[aria-label="${username}"]`);
      axios.post("http://127.0.0.1:3000/user/postime/new-postimer", {
        username: username
      }, {
        withCredentials: true,
        headers: {
          'Content-Type' : 'application/x-www-form-urlencoded; charset=UTF-8'
        }
      })
          .then(res => {
            if (res.data.check) {
              element.innerHTML = "followed"
            }
          })
          .catch(err => console.log(err))
    },
    getElapsedTime(t1) {
      let seconds = Math.floor((Date.now() - Date.parse(t1)) / 1000);
      if (seconds < 60) {
        return seconds + " s";
      }
      let mins = Math.floor(seconds / 60);
      if (mins < 60) {
        return mins + " m";
      }
      let hours = Math.floor(mins / 60);
      if (hours < 24) {
        return hours + " h";
      }
      let days = Math.floor(hours / 24);
      return days + " d";
    },
  },
};
</script>

<style scoped>
.main-find-postimer {
  width: 100%;
  text-align: center;
}
.container-card-find-postimer {
  display: inline-block;
}
.container-card-find-postimer > div {
  margin: 30px 0;
}
.card-find-postimer {
  background-color: #d9d9d9;
  display: flex;
  margin: 0 auto;
  justify-content: center;
  height: 145px;
  align-items: center;
  width: 600px;
}
.image-container-find-postimer img {
  margin-top: 5px;
  border-radius: 50%;
  width: 80px;
}
.verical-alingment-postimes-postimer-common {
  margin: 10px 20px;
}
.verical-alingment-postimes-postimer-common > div:first-child {
  margin: 0 0 20px 0;
}
.follow-btn {
  height: 30px;
  width: 130px;
  border: none;
  font-size: 16px;
  background: #12c6ff;
  border-radius: 20px;
  color: #000000;
}
</style>