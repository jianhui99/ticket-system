<template>
  <div id="app">
    <div id="main" role="main">
      <div class="top">
        <h1>You are now in line</h1>
        <span class="text-small text-bold">When it is your turn, you will have 5 minutes to enter the website.</span>

        <div class="queue-bar">
          <div class="queue-bar-fill"></div>
          <div class="queue-bar-fill"></div>
          <div class="queue-bar-fill"></div>
          <div class="queue-bar-fill"></div>
          <div class="queue-bar-fill"></div>
        </div>
      </div>
      <div class="middle">
        <p>There are 5 people in front of you.</p>

        <p>Estimated wait time: 5 minutes</p>

        <p>Status last updated: 11:23:45pm</p>
      </div>
      <div class="bottom">
        <p>Queue ID: 955dddd3-3aba-4bc8-af89-8649a626965d</p>
        <p>Thank you for your patience.</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
  #main {
    width: 100%;
    max-width: 400px;
    margin: 0 auto;
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 5px;
    background-color: #f9f9f9;
    text-align: left;
  }

  .top {
    margin-bottom: 20px;
  }

  .queue-bar {
    display: flex;
    justify-content: space-between;
    background-color: #ddd;
    border-radius: 5px;
    height: 20px;
  }

  .text-small {
    font-size: 12px;
  }

  .text-bold {
    font-weight: bold;
  }

  .queue-bar-fill {
    background-color: #4caf50;
    border-radius: 3px;
    height: 20px;
    width: 15%;
  }

  .middle {
    margin-bottom: 20px;
  }

  .bottom {
    font-size: 12px;
  }

  p {
    margin: 0;
  }

</style>

<script>
import axios from "axios";

export default {
  name: "App",
  data() {},
  mounted: function () {
    initUser();
  },
  methods: {},
};


function initUser() {
  axios.post("http://localhost:8080/api/v1/queues").then((response) => {
    console.log(response.data);
  });
}

</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
