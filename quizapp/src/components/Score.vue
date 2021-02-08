<template>
  <div class="row">
    <h4 class="score">Score: {{ score }}/{{ numberOfQuestions * 100 }}</h4>
  </div>
</template>

<script>
export default {
  name: "Score",
  props: ["score", "numberOfQuestions"],
};
</script>

<style scoped>
.score {
  margin-top: 5vh;
  margin-bottom: 2vh;
  text-align: center;
}
</style>
