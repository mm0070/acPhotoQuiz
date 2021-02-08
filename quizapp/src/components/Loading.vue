<template>
  <div class="container">
    <div class="row align-items-center loader"></div>
    <div class="row align-items-center loading-text">
      <h4>Loading questions: {{ loadingProgress }}/{{ numberOfQuestions }}</h4>
    </div>
  </div>
</template>

<script>
export default {
  name: "Loading",
  props: ["questions", "numberOfQuestions"],
  computed: {
    loadingProgress() {
      if (this.questions == undefined) {
        return "0";
      } else {
        return this.questions.length;
      }
    },
  },
};
</script>

<style scoped>
.loader,
.loader:before,
.loader:after {
  border-radius: 50%;
  width: 2.5em;
  height: 2.5em;
  -webkit-animation-fill-mode: both;
  animation-fill-mode: both;
  -webkit-animation: load7 1.8s infinite ease-in-out;
  animation: load7 1.8s infinite ease-in-out;
}
.loader {
  color: #0d6efd;
  font-size: 10px;
  margin: 80px auto;
  margin-top: 40vh;
  position: relative;
  text-indent: -9999em;
  -webkit-transform: translateZ(0);
  -ms-transform: translateZ(0);
  transform: translateZ(0);
  -webkit-animation-delay: -0.16s;
  animation-delay: -0.16s;
}
.loader:before,
.loader:after {
  content: "";
  position: absolute;
  top: 0;
}
.loader:before {
  left: -3.5em;
  -webkit-animation-delay: -0.32s;
  animation-delay: -0.32s;
}
.loader:after {
  left: 3.5em;
}
@-webkit-keyframes load7 {
  0%,
  80%,
  100% {
    box-shadow: 0 2.5em 0 -1.3em;
  }
  40% {
    box-shadow: 0 2.5em 0 0;
  }
}
@keyframes load7 {
  0%,
  80%,
  100% {
    box-shadow: 0 2.5em 0 -1.3em;
  }
  40% {
    box-shadow: 0 2.5em 0 0;
  }
}

.loading-text {
  text-align: center;
  margin-top: 10vh;
}
</style>
