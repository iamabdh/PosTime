<template>
  <div class="main-postimer-update">
    <div class="title-postimer-update">
      <h4>
        PosTimers Update
      </h4>
    </div>
    <div class="users-postimer-update-main">
      <div class="container-main-user-detail-postimer">
        <div class="user-postimer">
          <div class="user-image-postimer">
            <img src="https://avatars.githubusercontent.com/u/20264401">
          </div>
          <div class="user-details-postimer">
            <div class="time-latest-postimer">
              2 m
            </div>
            <div class="username-postimer">
              @jack
            </div>
          </div>
        </div>
      </div> <div class="container-main-user-detail-postimer">
        <div class="user-postimer">
          <div class="user-image-postimer">
            <img src="https://avatars.githubusercontent.com/u/20264401">
          </div>
          <div class="user-details-postimer">
            <div class="time-latest-postimer">
              2 m
            </div>
            <div class="username-postimer">
              @jack
            </div>
          </div>
        </div>
      </div> <div class="container-main-user-detail-postimer">
        <div class="user-postimer">
          <div class="user-image-postimer">
            <img src="https://avatars.githubusercontent.com/u/20264401">
          </div>
          <div class="user-details-postimer">
            <div class="time-latest-postimer">
              2 m
            </div>
            <div class="username-postimer">
              @jack
            </div>
          </div>
        </div>
      </div> <div class="container-main-user-detail-postimer">
        <div class="user-postimer">
          <div class="user-image-postimer">
            <img src="https://avatars.githubusercontent.com/u/20264401">
          </div>
          <div class="user-details-postimer">
            <div class="time-latest-postimer">
              2 m
            </div>
            <div class="username-postimer">
              @jack
            </div>
          </div>
        </div>
      </div> <div class="container-main-user-detail-postimer">
        <div class="user-postimer">
          <div class="user-image-postimer">
            <img src="https://avatars.githubusercontent.com/u/20264401">
          </div>
          <div class="user-details-postimer">
            <div class="time-latest-postimer">
              2 m
            </div>
            <div class="username-postimer">
              @jack
            </div>
          </div>
        </div>
      </div>
      </div>
  </div>
</template>

<script>
export default {
  name: "PostimersUpdateHome"
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