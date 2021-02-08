<template>
  <div class="fixed">
    <h4>Question {{ currentQuestionIndex + 1 }} of {{ numberOfQuestions }}</h4>
  </div>
</template>

<script>
export default {
  name: "Override",
  props: ["currentQuestionIndex", "numberOfQuestions"],
};
</script>

<style scoped>
div.fixed {
  position: fixed;
  top: 5px;
  left: 5px;
}
</style>
