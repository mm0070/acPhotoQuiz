<template>
  <div class="container">
    <div class="row align-items-center"></div>
    <div class="row align-items-center score-text">
      <h3>Your score: {{ score }}/1000</h3>
    </div>
  </div>
</template>

<script>
export default {
  name: "Finished",
  props: ["score"],
};
</script>

<style scoped>
.score-text {
  text-align: center;
  margin-top: 45vh;
}
</style>
