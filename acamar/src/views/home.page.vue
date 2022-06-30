<template>
  <div id="main-container">
    <div class="main-profile"></div>
    <div class="main-feed">
      <PosTimeCreate/>
      <PostimeViewHandlerComponent/>
    </div>
    <div class="main-update"></div>
  </div>
</template>

<script>
import axios from "axios";
import PosTimeCreate from "@/components/postime.create";
import PosTimeView from "@/components/postime.view";
import PostimeViewHandlerComponent from "@/components/postime.view.handler.component";

export default {
  name: "Home",
  components: {PostimeViewHandlerComponent, PosTimeView, PosTimeCreate},
  title: "Home",
  data() {
  },
  created() {
  },
  methods: {
  }
}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Poppins:wght@200&display=swap');
 * {
   font-family: 'Poppins', sans-serif;
 }
</style>