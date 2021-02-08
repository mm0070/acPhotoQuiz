<template>
  <p>your score: {{ score }}/1000</p>
</template>

<script>
export default {
  name: "Finished",
  props: ["score"]
};
</script>
