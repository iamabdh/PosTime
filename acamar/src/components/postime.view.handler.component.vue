<template>
  <PosTimeCreate/>
  <div>
    <div v-if="!postimes">
      no postime
    </div>
    <div class="postime-view-handler" v-if="postimes">
      <div v-for="postime in postimes">
        <PosTimeView v-bind:postime="postime" />
      </div>
    </div>
  </div>
</template>

<script>
import PosTimeCreate from "@/components/postime.create";
import PosTimeView from "@/components/postime.view";
import axios from "axios";

export default {
  name: "PostimeViewHandlerComponent",
  components: {PosTimeView,PosTimeCreate},
  data() {
    return {
      postimes: null
    }
  },
  created() {
    axios.get("http://127.0.0.1:3000/user/postime/feed-postimers", {
      withCredentials: true
    })
        .then(res => {
          this.postimes = res.data
        })
        .catch(err => console.log(err))
  }
}
</script>

<style scoped>

</style>