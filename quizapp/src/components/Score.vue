<template>
  <div>
    <p>Score: {{ score }}/{{ numberOfQuestions*100 }}</p>
  </div>
</template>

<script>
export default {
  name: "Score",
  props: ["score", "numberOfQuestions"],
};
</script>
