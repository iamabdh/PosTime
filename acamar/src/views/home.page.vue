<template>
  <div id="main-container">
    <div class="low-profile">
      <LowProfileHome @Change-Component-Feed="ChangeComponentFeed"/>
    </div>
    <div class="main-feed main-find-postimer">
      <component v-bind:is="componentHomeFeed"></component>
    </div>
    <div class="main-update">
      <PostimersUpdateHome />
    </div>
  </div>
</template>

<script>
import axios from "axios";
import PosTimeCreate from "@/components/postime.create";
import PosTimeView from "@/components/postime.view";
import PostimeViewHandlerComponent from "@/components/postime.view.handler.component";
import LowProfileHome from "@/components/low.profile.home";
import PostimersUpdateHome from "@/components/postimers.update.home";
import FindPosTimerHome from "@/components/find.postimer.home";
export default {
  name: "Home",
  components: {
    LowProfileHome,
    PostimeViewHandlerComponent,
    PosTimeCreate,
    PostimersUpdateHome,
    FindPosTimerHome
  },
  title: "Home",
  data() {
    return {
      componentHomeFeed: "PostimeViewHandlerComponent"
    }
  },
  created() {
  },
  methods: {
    ChangeComponentFeed(component) {
      this.componentHomeFeed = component
    }
  },
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Poppins:wght@200&display=swap');
 * {
   font-family: 'Poppins', sans-serif;
 }
#main-container {
  display: flex;
}
.low-profile {
  /*border: 1px solid;*/
  margin-right: auto;
  left: 0;
  width: 20%;
}

.main-feed {
  /*border: 1px solid;*/
  width: 50%;
}
.main-update {
  /*border: 1px solid;*/
  margin-left: auto;
  right: 0;
  width: 20%;
}
</style>