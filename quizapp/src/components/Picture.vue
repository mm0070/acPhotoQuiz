<template>
  <div>
    <img :src="imgSrc" />
  </div>
</template>

<script>
export default {
  name: "Picture",
  props: ["imgSrc"],
};
</script>

<style scoped>
img {
  display: block;
  margin-left: auto;
  margin-right: auto;
  margin-bottom: 1vh;
  max-height: 70vh;
  max-width: auto;
}
</style>
